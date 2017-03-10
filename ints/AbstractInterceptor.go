package ints

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type AbstractInterceptor struct {
    runtime.AbstractComponent
    
    logic runtime.IController
}

func NewAbstractInterceptor(id string, config *runtime.DynamicMap) *AbstractInterceptor {
    return &AbstractInterceptor { AbstractComponent: *runtime.NewAbstractComponent(id, config) }    
}

func (c *AbstractInterceptor) Logic() runtime.IController {
    return c.logic
}

func (c *AbstractInterceptor) Init(refs *runtime.References) error {
    err := c.AbstractComponent.Init(refs)
    if err != nil { return err }
        
    c.logic = refs.Logic
    if c.logic == nil { 
       c.NewInternalError("Controller is not specified") 
    }
        
    return nil
}

