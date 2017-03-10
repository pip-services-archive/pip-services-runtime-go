package db

import (
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-services-runtime-go/util"
)

type AbstractDataAccess struct {
    runtime.AbstractComponent
}

func NewAbstractDataAccess(id string, config *runtime.DynamicMap) *AbstractDataAccess {
    return &AbstractDataAccess { AbstractComponent: *runtime.NewAbstractComponent(id, config) }    
}

func (c *AbstractDataAccess) CreateUUID() string {
    return util.IdGenerator.UUID()
}

func (c *AbstractDataAccess) Clear() {}