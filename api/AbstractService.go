package api

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type AbstractService struct {
    runtime.AbstractComponent
    
    logic runtime.IController
}

func NewAbstractService(id string, config *runtime.DynamicMap) *AbstractService {
    return &AbstractService { AbstractComponent: *runtime.NewAbstractComponent(id, config) }    
}

func (c *AbstractService) Logic() runtime.IController {
    return c.logic
}

func (c *AbstractService) Init(refs *runtime.References) error {
    err := c.AbstractComponent.Init(refs)
    if err != nil { return err }

    c.logic = refs.Logic
    
    if c.logic == nil {
        return c.NewInternalError("Controller is not specified") 
    }
    
    return nil
}
