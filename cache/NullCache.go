package cache

import (
    "github.com/pip-services/pip-services-runtime-go"    
)

type NullCache struct {
    AbstractCache
}

func NewNullCache(config *runtime.DynamicMap) *NullCache {
    return &NullCache { AbstractCache: *NewAbstractCache("NullCache", config) }
}

func (c *NullCache) Get(key string, value interface{}) (interface{}, error) {
    return nil, nil
}

func (c *NullCache) Set(key string, value interface{}) (interface{}, error) {
    return value, nil
}

func (c *NullCache) Unset(key string) error {
    return nil
}
