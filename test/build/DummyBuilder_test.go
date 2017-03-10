package build

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/test/db"    
)

func TestBuilderStart(t *testing.T) {
    builder := NewDummyBuilder()
    
    err := builder.Start(runtime.NewEmptyMap())
    assert.Nil(t, err)
    
    assert.NotNil(t, builder.Config())
    assert.Len(t, builder.References().GetComponents(), 0)
    assert.NotNil(t, builder.References())
}

func TestBuilderStartWithConfig(t *testing.T) {
    buildConfig := runtime.NewMapAndSet(
		"config.type", "direct",
		"log.type", "console",
		"counters.type", "log",
		"cache.type", "memory",
		"db.type", "file",
		"db.path", "dummies.json",
		"db.data", []db.Dummy{},
		"api.type", "rest",
		"api.version", 1 )
    
    builder := NewDummyBuilder()
    
    err := builder.StartWithConfig(buildConfig)
    assert.Nil(t, err)
    assert.NotNil(t, builder.Config())
}

func TestBuilderStartWithFile(t *testing.T) {
    builder := NewDummyBuilder()
    
    err := builder.StartWithFile("./config.json")
    assert.Nil(t, err)
    assert.NotNil(t, builder.Config())
}

func TestBuilderDefaults(t *testing.T) {
    builder := NewDummyBuilder()
    
    err := builder.StartWithConfig(nil)
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 0)
    
    err = builder.BuildLog()
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 1)
    assert.NotNil(t, builder.References().Log)
    
    err = builder.BuildCounters()
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 2)
    assert.NotNil(t, builder.References().Counters)

    err = builder.BuildCache()
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 3)
    assert.NotNil(t, builder.References().Cache)
    
    err = builder.BuildDataAccess()
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 3)
    assert.Nil(t, builder.References().DB)
    
    err = builder.BuildDependencies()
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 3)
    assert.Len(t, builder.References().Deps, 0)
    
    err = builder.BuildController()
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 4)
    assert.NotNil(t, builder.References().Ctrl)

    err = builder.BuildInterceptors()
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 4)
    assert.Len(t, builder.References().Ints, 0)

    err = builder.BuildApi()
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 4)
    assert.Len(t, builder.References().API, 0)

    err = builder.BuildAddons()
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 4)
    assert.Len(t, builder.References().Addons, 0)
}

func TestBuilderBuildAll(t *testing.T) {
    builder := NewDummyBuilder()
    
    err := builder.StartWithFile("./config.json")
    assert.Len(t, builder.References().GetComponents(), 0)
    
    err = builder.BuildAll()
    assert.Nil(t, err)
    assert.Len(t, builder.References().GetComponents(), 7)
    assert.NotNil(t, builder.References().Log)
    assert.NotNil(t, builder.References().Counters)
    assert.NotNil(t, builder.References().Cache)
    assert.NotNil(t, builder.References().DB)
    assert.Len(t, builder.References().Deps, 1)
    assert.NotNil(t, builder.References().Ctrl)
    assert.Len(t, builder.References().Ints, 0)
    assert.Len(t, builder.References().API, 1)
    assert.Len(t, builder.References().Addons, 0)
}
