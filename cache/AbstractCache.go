package  cache

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type AbstractCache struct {
    runtime.AbstractComponent
}

func NewAbstractCache(id string, config *runtime.DynamicMap) *AbstractCache {
    return &AbstractCache { AbstractComponent: *runtime.NewAbstractComponent(id, config) }    
}
