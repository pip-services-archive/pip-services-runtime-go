package discovery

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type AbstractDiscovery struct {
    runtime.AbstractComponent
}

func NewAbstractDiscovery(id string, config *runtime.DynamicMap) *AbstractDiscovery {
    return &AbstractDiscovery { AbstractComponent: *runtime.NewAbstractComponent(id, config) }    
}
