package deps

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type AbstractClient struct {
    runtime.AbstractComponent
}

func NewAbstractClient(id string, config *runtime.DynamicMap) *AbstractClient {
    return &AbstractClient { AbstractComponent: *runtime.NewAbstractComponent(id, config) }    
}

func (c *AbstractClient) Instrument(name string) *runtime.Timing {
    c.Trace("Calling " + name + " method")
    return c.BeginTiming(name + ".CallTime")
}

