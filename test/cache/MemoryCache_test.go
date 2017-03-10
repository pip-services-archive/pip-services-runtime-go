package cache

import (
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/cache"    
)

type MemoryCacheTest struct {
    suite.Suite
    cache runtime.ICache
    fixture *CacheFixture
}

func (suite *MemoryCacheTest) SetupTest() {
    refs := runtime.NewReferences()
    
    config := runtime.NewMapAndSet(
        "timeout", 500,
    )
    
    suite.cache = cache.NewMemoryCache(config)
    suite.cache.Init(refs)
    suite.cache.Open()
    
    suite.fixture = NewCacheFixture(suite.cache)
}

func (suite *MemoryCacheTest) TearDownTest() {
    suite.cache.Close()
}

func (suite *MemoryCacheTest) TestBasicOperations() {
    suite.fixture.TestBasicOperations(suite.T())
}

func (suite *MemoryCacheTest) TestReadAfterTimeout() {
    value, err := suite.cache.Set("test", 123)
    assert.Equal(suite.T(), 123, value)
    assert.Nil(suite.T(), err)
    
    value, err = suite.cache.Get("test", nil)
    assert.Equal(suite.T(), 123, value)
    assert.Nil(suite.T(), err)
    
    time.Sleep(1000 * time.Millisecond)
    value, err = suite.cache.Get("test", nil)
    assert.Nil(suite.T(), value)
    assert.Nil(suite.T(), err)
}

func TestMemoryCacheTestSuite(t *testing.T) {
    suite.Run(t, new(MemoryCacheTest))
}