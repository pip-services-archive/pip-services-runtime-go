package config

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type AbstractConfig struct {
    runtime.AbstractComponent
}

func NewAbstractConfig(id string, config *runtime.DynamicMap) *AbstractConfig {
    return &AbstractConfig { AbstractComponent: *runtime.NewAbstractComponent(id, config) }
}