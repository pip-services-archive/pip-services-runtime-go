package addons

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type AbstractAddon struct {
    runtime.AbstractComponent
}

func NewAbstractAddon(id string, config *runtime.DynamicMap) *AbstractAddon {
    return &AbstractAddon { AbstractComponent: *runtime.NewAbstractComponent(id, config) }    
}
