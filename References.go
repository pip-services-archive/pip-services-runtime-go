package runtime

type References struct {
    Config      IConfig
    Log         ILog
    Counters    ICounters
    Cache       ICache
    DB          IDataAccess
    Deps        map[string]IClient
    Ctrl        IController
    Ints        []IInterceptor
    Logic       IController
    API         []IService
    Addons      []IAddon
}

func NewReferences() *References {
    return &References{
        Config: nil,
        Log: nil,
        Counters: nil,
        Cache: nil,
        DB: nil,
        Deps: map[string]IClient{},
        Ctrl: nil,
        Ints: []IInterceptor{},
        Logic: nil,
        API: []IService{},
        Addons: []IAddon{}, 
    }
}

func (refs *References) GetDep(name string) IClient {
    if (refs.Deps != nil) { return refs.Deps[name]; }
    return nil;
}

func (refs *References) GetComponents() []IComponent {
    r := []IComponent {}
    
    if (refs.Config != nil) { r = append(r, refs.Config) }
    if (refs.Log != nil) { r = append(r, refs.Log) }
    if (refs.Counters != nil) { r = append(r, refs.Counters) }
    if (refs.Cache != nil) { r = append(r, refs.Cache) }
    if (refs.DB != nil) { r = append(r, refs.DB) }
    if (refs.Deps != nil) {
        for _, dep := range refs.Deps { r = append(r, dep) } 
    }
    if (refs.Ctrl != nil) { r = append(r, refs.Ctrl) }
    if (refs.Ints != nil) {
        for _, ints := range refs.Ints { r = append(r, ints) } 
    }
    if (refs.API != nil) {
        for _, api := range refs.API { r = append(r, api) } 
    }
    if (refs.Addons != nil) {
        for _, addon := range refs.Addons { r = append(r, addon) } 
    }
    
    return r
}

func (refs *References) WithConfig(config IConfig) *References {
    refs.Config = config
    return refs
}

func (refs *References) WithLog(log ILog) *References {
    refs.Log = log
    return refs
}

func (refs *References) WithCounters(counters ICounters) *References {
    refs.Counters = counters
    return refs
}

func (refs *References) WithCache(cache ICache) *References {
    refs.Cache = cache
    return refs
}

func (refs *References) WithDB(db IDataAccess) *References {
    refs.DB = db
    return refs
}

func (refs *References) WithDep(name string, dep IClient) *References {
    refs.Deps[name] = dep
    return refs
}

func (refs *References) WithDeps(params ...interface{}) *References {
    deps := map[string]IClient{}
    
    for i := 0; i < len(params) - 1; i += 2 {
        name := params[i].(string)
        dep := params[i + 1].(IClient)
        deps[name] = dep
    }
    
    refs.Deps = deps
    return refs
}

func (refs *References) WithCtrl(ctrl IController) *References {
    refs.Ctrl = ctrl
    return refs
}

func (refs *References) WithInts(ints ...IInterceptor) *References {
    refs.Ints = ints
    return refs
}

func (refs *References) WithLogic(logic IController) *References {
    refs.Logic = logic
    return refs
}

func (refs *References) WithAPI(api ...IService) *References {
    refs.API = api
    return refs
}

func (refs *References) WithAddons(addons ...IAddon) *References {
    refs.Addons = addons
    return refs
}