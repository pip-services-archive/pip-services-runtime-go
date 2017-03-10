package logic

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type AbstractController struct {
    runtime.AbstractComponent
    
    db runtime.IDataAccess
    deps map[string]runtime.IClient
}

func NewAbstractController(id string, config *runtime.DynamicMap) *AbstractController {
    return &AbstractController { AbstractComponent: *runtime.NewAbstractComponent(id, config) }    
}

func (c *AbstractController) DB() runtime.IDataAccess {
    return c.db
}

func (c *AbstractController) Deps() map[string]runtime.IClient {
    return c.deps
}

func (c *AbstractController) Init(refs *runtime.References) error {
    err := c.AbstractComponent.Init(refs)
    if err != nil { return err }
        
    c.db = refs.DB
    c.deps = refs.Deps
    
    return nil
}

func (c *AbstractController) Instrument(name string) *runtime.Timing {
    c.Trace("Executing " + name + " method")
    return c.BeginTiming(name + ".ExecTime")
}

