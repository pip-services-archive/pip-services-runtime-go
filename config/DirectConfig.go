package config

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type DirectConfig struct {
    AbstractConfig
}

func NewDirectConfig(config *runtime.DynamicMap) *DirectConfig {
    return &DirectConfig { AbstractConfig: *NewAbstractConfig("DirectConfig", config) }
}

func (c *DirectConfig) Read() (*runtime.DynamicMap, error) {
    return c.Config(), nil
}
