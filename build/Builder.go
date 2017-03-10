package build

import (
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-services-runtime-go/log"
    "github.com/pip-services/pip-services-runtime-go/counters"
    "github.com/pip-services/pip-services-runtime-go/cache"
    "github.com/pip-services/pip-services-runtime-go/config"
)

type Builder struct {
    name string
    types *runtime.DynamicMap
    config *runtime.DynamicMap
    refs *runtime.References
}

func NewBuilder(types *runtime.DynamicMap, name string) *Builder {
    defaultTypes := runtime.NewMapAndSet(
        "log.null", func(c *runtime.DynamicMap) runtime.IComponent { return log.NewNullLog(c) },
        "log.console", func(c *runtime.DynamicMap) runtime.IComponent { return log.NewConsoleLog(c) },
        "counters.null", func(c *runtime.DynamicMap) runtime.IComponent { return counters.NewNullCounters(c) },
        "counters.log", func(c *runtime.DynamicMap) runtime.IComponent { return counters.NewLogCounters(c) },
        "cache.null", func(c *runtime.DynamicMap) runtime.IComponent { return cache.NewNullCache(c) },
        "cache.memory", func(c *runtime.DynamicMap) runtime.IComponent { return cache.NewMemoryCache(c) },
        "config.direct", func(c *runtime.DynamicMap) runtime.IComponent { return config.NewDirectConfig(c) },
        "config.json", func(c *runtime.DynamicMap) runtime.IComponent { return config.NewJsonConfig(c) },
        "config.file", func(c *runtime.DynamicMap) runtime.IComponent { return config.NewJsonConfig(c) } )
    types = runtime.NewMapWithDefaults(types, defaultTypes)
    
    c := Builder { name: name, types: types }
    c.refs = runtime.NewReferences()
    
    return &c
}

func (c *Builder) Name() string { 
    return c.name
}

func (c *Builder) Config() *runtime.DynamicMap {
    return c.config
}

func (c *Builder) References() *runtime.References {
    return c.refs
}

func (c *Builder) newBuildError(message string) *runtime.MicroserviceError {
    return runtime.NewBuildError(message).ForComponent(c.name)
}

func (c *Builder) newStateError(message string) *runtime.MicroserviceError {
    return runtime.NewStateError(message).ForComponent(c.name)
}

func (c *Builder) newConfigError(message string) *runtime.MicroserviceError {
    return runtime.NewConfigError(message).ForComponent(c.name)
}

func (c *Builder) instantiateComponent(id string, creator interface{}, config *runtime.DynamicMap) (runtime.IComponent, error) {
    var creatorFunc, ok = creator.(func(config *runtime.DynamicMap) runtime.IComponent) 
    if ok == false {
        return nil, c.newConfigError(id + " type is not correctly configured") 
    }

    component := creatorFunc(config)
    if component == nil {
        return nil, c.newBuildError("Failed to instantiate " + id + " component")
    }
    return component, nil
} 

func (c *Builder) createCustomComponent(id string, types *runtime.DynamicMap, config *runtime.DynamicMap) (runtime.IComponent, error) {
    return nil, c.newConfigError("Custom components are not supported by GO runtime")
} 

func (c *Builder) createStandardComponent(id string, types *runtime.DynamicMap, config *runtime.DynamicMap) (runtime.IComponent, error) {
    typeName := config.GetString("type")
    
    creator := types.Get(typeName)
    if creator == nil {
        return nil, c.newConfigError(id + " type is not correctly configured").WithDetails(typeName)
    }
        
    return c.instantiateComponent(id, creator, config)
}

func (c *Builder) createComponent(id string, types *runtime.DynamicMap, config *runtime.DynamicMap) (runtime.IComponent, error) {
    typ := config.GetString("type")
    if typ == "" {
        return nil, c.newConfigError(id + " type is not configured")
    }
    
    var component runtime.IComponent
    var err error    
    if typ == "custom" { 
        component, err = c.createCustomComponent(id, types, config) 
    } else { 
        component, err = c.createStandardComponent(id, types, config) 
    }         
    if err != nil { return nil, err }
    
    if component == nil {
        return nil, c.newBuildError(id + " class does not implement IComponent interface").WithDetails(typ)
    }
    return component, nil
}

func (c *Builder) checkStarted() error {
    if c.config == nil {
        return c.newStateError("Building process was not started")
    }
    return nil
}

func (c *Builder) sectionConfig(config *runtime.DynamicMap, section string, defaultConfig *runtime.DynamicMap) *runtime.DynamicMap {
    // Extract section configuration
    result := config.GetNullableMap(section)
    
    if result == nil { result = defaultConfig }
    
    return result
}

func (c *Builder) sectionConfigs(config *runtime.DynamicMap, section string, defaultConfig *runtime.DynamicMap) []*runtime.DynamicMap {    
    defaultResult := []*runtime.DynamicMap {}
    if defaultConfig != nil { defaultResult = append(defaultResult, defaultConfig) }
    
    // Extract section configuration
    sectionConfig := config.Get(section)
    if sectionConfig == nil { return defaultResult }
    
    // Try to get array of configs 
    sectionArray, ok := sectionConfig.([]interface{})    
    // If config is single then return
    result := []*runtime.DynamicMap {}
    if ok == false {
        result := append(result, runtime.NewMapOf(sectionConfig)) 
        return result 
    }

    // Copy over section configs
    for _, p := range sectionArray {
        result = append(result, runtime.NewMapOf(p))
    }
    
    // Return configs or default if empty
    if len(result) == 0 { return defaultResult }
    return result
}

func (c *Builder) Start(config *runtime.DynamicMap) error {
    // Initialize references
    c.refs = runtime.NewReferences()
    
    // Create configuration
    defaultConfig := runtime.NewMapAndSet("config.type", "direct")
    config = runtime.NewMapWithDefaults(config, defaultConfig)
    configForConfig := config.GetMap("config")
    
    // Get config types
    configTypes := c.types.GetMap("config")

    // Construct config reader    
    component, err := c.createComponent("Config", configTypes, configForConfig)
    if err != nil { return err }
    reader, ok := component.(runtime.IConfig)
    if ok == false {
        return c.newBuildError("Config component does not implement IConfig interface")
    }

    // Initialize and open config reader    
    reader.Init(runtime.NewReferences())
    err = reader.Open()
    if err != nil { return err }
    
    // Read configuration
    c.config, err = reader.Read()
    if err != nil { return err }
    
    // Set default empty configuration
    if c.config == nil { c.config = runtime.NewEmptyMap() }
    
    return nil
}

func (c *Builder) StartWithConfig(config *runtime.DynamicMap) error {
    // Initialize references
    c.refs = runtime.NewReferences()
    c.config = config
    
    // Set default empty configuration
    if c.config == nil { c.config = runtime.NewEmptyMap() }
    
    return nil
}

func (c *Builder) StartWithFile(path string) error {
    // Initialize references
    c.refs = runtime.NewReferences()
    
    // Read configuration from file
    var err error
    c.config, err = config.ReadConfig(path)
    if err != nil { return err }

    // Set default empty configuration
    if c.config == nil { c.config = runtime.NewEmptyMap() }

    return nil
}

func (c *Builder) SetReferences(refs *runtime.References) error {
    err := c.checkStarted()
    if err != nil { return err }
    
    c.refs = refs
    return nil
}

func (c *Builder) BuildDiscovery() error {
    err := c.checkStarted()
    if err != nil { return err }
    
    // // Define discovery types and configuration
    // types := c.types.GetMap("discovery")
    // defaultConfig := runtime.NewMapAndSet("type", "null")
    // config := c.sectionConfig(c.config, "discovery", defaultConfig)

    // // Create discovery component
    // component, err := c.createComponent("Discovery", types, config)
    // if err != nil { return err }
    
    // discovery, ok := component.(runtime.IDiscovery)
    // if ok == false {
    //     return c.newBuildError("Discovery class does not implement IDiscovery interface").WithDetails(config)
    // }
    
    // c.refs.Discovery = discovery
    
    return nil
}

func (c *Builder) BuildLog() error {
    err := c.checkStarted()
    if err != nil { return err }

    // Define log types and configurations    
    types := c.types.GetMap("log")
    defaultConfig := runtime.NewMapAndSet("type", "null")
    configs := c.sectionConfigs(c.config, "log", defaultConfig)
    
    // Create log components
    logs := make([]runtime.ILog, 0, len(configs))
    for _, config := range configs {
        component, err := c.createComponent("Log", types, config)
        if err != nil { return err }
        
        log, ok := component.(runtime.ILog)
        if ok == false {
            return c.newBuildError("Log class does not implement ILog interface").WithDetails(config)
        }
        
        logs = append(logs, log)
    }

    // Define log    
    if len(logs) == 1 {
        c.refs.Log = logs[0]
    } else {
        compositeLog := log.NewCompositeLog(logs)
        c.refs.Log = compositeLog
    }
    
    return nil
}

func (c *Builder) BuildCounters() error {
    err := c.checkStarted()
    if err != nil { return err }
    
    // Define counters types and configuration
    types := c.types.GetMap("counters")
    defaultConfig := runtime.NewMapAndSet("type", "null")
    config := c.sectionConfig(c.config, "counters", defaultConfig)

    // Create counters component
    component, err := c.createComponent("Counters", types, config)
    if err != nil { return err }
    
    counters, ok := component.(runtime.ICounters)
    if ok == false {
        return c.newBuildError("Counters class does not implement ICounters interface").WithDetails(config)
    }
    
    c.refs.Counters = counters
    
    return nil
}

func (c *Builder) BuildCache() error {
    err := c.checkStarted()
    if err != nil { return err }
    
    // Define cache types and configuration
    types := c.types.GetMap("cache")
    defaultConfig := runtime.NewMapAndSet("type", "null")
    config := c.sectionConfig(c.config, "cache", defaultConfig)

    // Create cache component
    component, err := c.createComponent("Cache", types, config)
    if err != nil { return err }
    
    cache, ok := component.(runtime.ICache)
    if ok == false {
        return c.newBuildError("Cache class does not implement ICache interface").WithDetails(config)
    }
    
    c.refs.Cache = cache
    
    return nil
}

func (c *Builder) BuildDataAccess() error {
    err := c.checkStarted()
    if err != nil { return err }
    
    // Define db types and configuration
    types := c.types.GetMap("db")
    config := c.sectionConfig(c.config, "db", nil)
    
    // Check data access configuration
    if config == nil && c.config.Has("db") {
        return c.newBuildError("Incorrect Db section configuration").WithDetails(config)
    }
    
    // Process empty data access
    if config == nil {
        c.refs.DB = nil
        return nil
    }
    
    // Create db component
    component, err := c.createComponent("Db", types, config)
    if err != nil { return err }
    
    db, ok := component.(runtime.IDataAccess)
    if ok == false {
        return c.newBuildError("Db class does not implement IDataAccess interface").WithDetails(config)
    }
    
    c.refs.DB = db
    
    return nil
}

func (c *Builder) BuildDependencies() error {
    err := c.checkStarted()
    if err != nil { return err }
    
    // Define deps types and configuration
    types := c.types.GetMap("deps")
    configs := c.sectionConfigs(c.config, "deps", nil)
    
    // Create deps components
    for _, config := range configs {
        name := config.GetString("name")
        if name == "" {
            return c.newBuildError("Dep has no specified name").WithDetails(config)
        }
        
        nameTypes := types.GetMap(name)
        component, err := c.createComponent("Dep", nameTypes, config)
        if err != nil { return err }
        
        dep, ok := component.(runtime.IClient)
        if ok == false {
            return c.newBuildError("Dep class does not implement IClient interface").WithDetails(config)
        }
        
        c.refs.Deps[name] = dep
    } 
    
    return nil
}

func (c *Builder) BuildController() error {
    err := c.checkStarted()
    if err != nil { return err }
    
    // Define ctrl types and configuration
    types := c.types.GetMap("ctrl")
    config := c.sectionConfig(c.config, "ctrl", nil)
    defaultConfig := runtime.NewMapAndSet("type", "default")
    config = runtime.NewMapWithDefaults(config, defaultConfig)
    
    // Create ctrl component
    component, err := c.createComponent("Ctrl", types, config)
    if err != nil { return err }
    
    ctrl, ok := component.(runtime.IController)
    if ok == false {
        c.newBuildError("Ctrl class does not implement IController interface").WithDetails(config)
    }
    
    c.refs.Ctrl = ctrl
   
    return nil
}

func (c *Builder) BuildInterceptors() error {
    err := c.checkStarted()
    if err != nil { return err }
    
    // Define ints types and configuration
    types := c.types.GetMap("ints")
    configs := c.sectionConfigs(c.config, "ints", nil)
    
    // Create ints components
    for _, config := range configs {
        component, err := c.createComponent("Int", types, config)
        if err != nil { return err }
        
        ints, ok := component.(runtime.IInterceptor)
        if ok == false {
            return c.newBuildError("Int class does not implement IInterceptor interface").WithDetails(config)
        }
        
        c.refs.Ints = append(c.refs.Ints, ints)
    } 
    
    return nil
}

func (c *Builder) BuildApi() error {
    err := c.checkStarted()
    if err != nil { return err }
    
    // Define api types and configuration
    types := c.types.GetMap("api")
    configs := c.sectionConfigs(c.config, "api", nil)
    
    // Create api components
    for _, config := range configs {
        version := "version" + config.GetString("version")
        if version == "version" { version = "default" }

        versionTypes := types.GetMap(version)
        if len(*versionTypes) == 0 {
            return c.newConfigError("Unsupported API version").WithDetails(version)
        }
        
        component, err := c.createComponent("Api", versionTypes, config)
        if err != nil { return err }
        
        api, ok := component.(runtime.IService)
        if ok == false {
            return c.newBuildError("Api class does not implement IService interface").WithDetails(config)
        }
        
        c.refs.API = append(c.refs.API, api)
    } 

    // if len(c.refs.API) == 0 {        
    //     return c.newConfigError("At least one API service must be configured")
    // }
    
    return nil
}

func (c *Builder) BuildAddons() error {
    err := c.checkStarted()
    if err != nil { return err }
    
    // Define addons types and configuration
    types := c.types.GetMap("addons")
    configs := c.sectionConfigs(c.config, "addons", nil)
    
    // Create addons components
    for _, config := range configs {
        component, err := c.createComponent("Addon", types, config)
        if err != nil { return err }
        
        addon, ok := component.(runtime.IAddon)
        if ok == false {
            return c.newBuildError("Addon class does not implement IAddon interface").WithDetails(config)
        }
        
        c.refs.Addons = append(c.refs.Addons, addon)
    } 
    
    return nil
}

func (c *Builder) BuildAll() error {
    err := c.checkStarted()
    if err != nil { return err }
    
    err = c.BuildDiscovery()
    if err != nil { return err }

    err = c.BuildLog()
    if err != nil { return err }

    err = c.BuildCounters()
    if err != nil { return err }

    err = c.BuildCache()
    if err != nil { return err }

    err = c.BuildDataAccess()
    if err != nil { return err }

    err = c.BuildDependencies()
    if err != nil { return err }

    err = c.BuildController()
    if err != nil { return err }

    err = c.BuildInterceptors()
    if err != nil { return err }

    err = c.BuildApi()
    if err != nil { return err }

    err = c.BuildAddons()
    if err != nil { return err }
    
    return nil
}
