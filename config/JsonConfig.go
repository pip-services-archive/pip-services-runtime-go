package config

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "github.com/pip-services/pip-services-runtime-go"
)

type JsonConfig struct {
    AbstractConfig
}

func NewJsonConfig(config *runtime.DynamicMap) *JsonConfig {
    return &JsonConfig { AbstractConfig: *NewAbstractConfig("JsonConfig", config) }
}

func NewJsonConfigFromFile(path string) *JsonConfig {
    config := &runtime.DynamicMap { "path": path }
    return NewJsonConfig(config)
}

func (c *JsonConfig) Init(refs *runtime.References) error {
    if c.Config().HasNot("path") { 
        c.NewConfigError("Config file path is missing in configuration") 
    }
    
    return c.AbstractConfig.Init(refs)
}

func (c *JsonConfig) Open() error {
    path := c.Config().GetString("path")
    
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return c.NewReadError("Config file was not found", err).WithDetails(path)
    }
    
    return c.AbstractConfig.Open()
}

func (c *JsonConfig) Read() (*runtime.DynamicMap, error) {
    path := c.Config().GetString("path")
    return ReadConfig(path)
}

func ReadConfig(path string) (*runtime.DynamicMap, error) {
    if path == "" { 
        return nil, runtime.NewConfigError("Missing config file path").ForComponent("JsonConfig")
    }
    
    raw, err := ioutil.ReadFile(path)
    if err != nil { 
        return nil, runtime.NewReadError("Failed reading configuration", err).ForComponent("JsonConfig")
    }
    
    value := map[string] interface{} {}
    err = json.Unmarshal(raw, &value)
    if err != nil { 
        return nil, runtime.NewReadError("Failed to unmarshal config", err).ForComponent("JsonConfig")
    }
    
    return runtime.NewMapOf(value), nil
}