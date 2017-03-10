package build

import (
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/build"    
    "github.com/pip-services/pip-services-runtime-go/test/db"    
    "github.com/pip-services/pip-services-runtime-go/test/logic"    
    "github.com/pip-services/pip-services-runtime-go/test/deps"    
    "github.com/pip-services/pip-services-runtime-go/test/api"    
)

func NewDummyBuilder() *build.Builder {
    types := runtime.NewMapAndSet(
        "db.file", func(c *runtime.DynamicMap) runtime.IComponent { return db.NewDummyFileDataAccess(c) },
        "deps.dummy.rest", func(c *runtime.DynamicMap) runtime.IComponent { return deps.NewDummyRestClient(c) },
        "ctrl.default", func(c *runtime.DynamicMap) runtime.IComponent { return logic.NewDummyController(c) },
        "api.version1.rest", func(c *runtime.DynamicMap) runtime.IComponent { return api.NewDummyRestService(c) },
        "api.default.rest", func(c *runtime.DynamicMap) runtime.IComponent { return api.NewDummyRestService(c) } )
    return build.NewBuilder(types, "Dummy.Builder")
}