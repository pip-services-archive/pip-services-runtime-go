package cache

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go"    
)

type CacheFixture struct {
    Cache runtime.ICache
}

func NewCacheFixture(Cache runtime.ICache) *CacheFixture {
    return &CacheFixture{ Cache: Cache }
}

func (c *CacheFixture) TestBasicOperations(t *testing.T) {
    // Set value
    value, err := c.Cache.Set("test", 123)
    assert.Equal(t, 123, value)
    assert.Nil(t, err)
    
    value, err = c.Cache.Get("test", value)
    assert.Equal(t, 123, value)
    assert.Nil(t, err)

    // Set null
    value, err = c.Cache.Set("test", nil)
    assert.Nil(t, value)
    assert.Nil(t, err)

    value, err = c.Cache.Get("test", value)
    assert.Nil(t, value)
    assert.Nil(t, err)

    // Set another value
    value, err = c.Cache.Set("test", "ABC")
    assert.Equal(t, "ABC", value)
    assert.Nil(t, err)
    
    value, err = c.Cache.Get("test", value)
    assert.Equal(t, "ABC", value)
    assert.Nil(t, err)

    // Unset value    
    err = c.Cache.Unset("test")
    assert.Nil(t, err)

    value, err = c.Cache.Get("test", value)
    assert.Nil(t, value)
    assert.Nil(t, err)
}

