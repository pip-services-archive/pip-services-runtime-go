package run

import (
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-services-runtime-go/build"
    "github.com/pip-services/pip-services-runtime-go/config"
)

type Microservice struct {
    builder *build.Builder
    config *runtime.DynamicMap
    refs *runtime.References
    components []runtime.IComponent
}

func NewMicroservice(builder *build.Builder) *Microservice {
    return &Microservice { builder: builder }
}

func NewMicroserviceWithConfig(builder *build.Builder, config *runtime.DynamicMap) *Microservice {
    return &Microservice { builder: builder, config: config }
}

func (c *Microservice) SetConfig(config interface{}) {
    c.config = runtime.NewMapOf(config)
}

func (c *Microservice) LoadConfig(configPath string) error {
    var err error
    c.config, err = config.ReadConfig(configPath)
    return err
}

func (c *Microservice) Fatal(message interface{}) {
    build.LifeCycleManager.Fatal(c.refs, message)
}

func (c *Microservice) Error(message interface{}) {
    build.LifeCycleManager.Error(c.refs, message)
}

func (c *Microservice) Info(message interface{}) {
    build.LifeCycleManager.Info(c.refs, message)
}

func (c *Microservice) Trace(message interface{}) {
    build.LifeCycleManager.Trace(c.refs, message)
}

func (c *Microservice) build() error {
    if c.config == nil { 
        return runtime.NewInternalError("Microservice configuration was not set"); 
    }
    
    err := c.builder.StartWithConfig(c.config)
    if err != nil { return err }
    
    err = c.builder.BuildAll()
    if err != nil { return err }
    
    c.refs = c.builder.References()
    c.components = c.builder.Components()
    
    return nil
}

func (c *Microservice) init() error {
    c.Trace("Initializing " + c.builder.Name() + " microservice")
    return build.LifeCycleManager.Init(c.refs, c.components)
}

func (c *Microservice) open() error {
    c.Trace("Opening " + c.builder.Name() + " microservice")
    err := build.LifeCycleManager.Open(c.refs, c.components)
    if err != nil { return err }

    c.Info("Microservice " + c.builder.Name() + " started")
    return nil
}

func (c *Microservice) Start() error {
    if err := c.build(); err != nil { return err }
    if err := c.init(); err != nil { return err }
    if err := c.open(); err != nil { return err }
    return nil
}

func (c *Microservice) StartWithConfig(config *runtime.DynamicMap) error {
    c.SetConfig(config)
    return c.Start()       
}

func (c *Microservice) StartWithConfigFile(configPath string) error {
    if err := c.LoadConfig(configPath); err != nil { return err }
    return c.Start()       
}

func (c *Microservice) Stop() error {
    c.Trace("Closing " + c.builder.Name() + " microservice")
    err := build.LifeCycleManager.ForceClose(c.refs, c.components)
    c.Info("Microservice " + c.builder.Name() + " stopped")
    return err
}
