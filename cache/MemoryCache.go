// Todo:
// - Add performance counters
// - Add tracing
// - Synchronize for concurrent access

package cache

import (
    "time"
    "github.com/pip-services/pip-services-runtime-go"    
)

type MemoryCache struct {
    AbstractCache
    cache   map[string]*CacheEntry
    timeout time.Duration
    maxSize int
    count int
}

func NewMemoryCache(config *runtime.DynamicMap) *MemoryCache {
    defaultConfig := runtime.NewMapAndSet(
        "timeout", 60000, // timeout in milliseconds
        "maxSize", 1000, // maximum number of elements in cache
    )
    config = runtime.NewMapWithDefaults(config, defaultConfig)

    c := &MemoryCache { AbstractCache: *NewAbstractCache("MemoryCache", config) }
    c.cache = map[string]*CacheEntry {}
    c.timeout = time.Duration(config.GetLong("timeout")) * time.Millisecond
    c.maxSize = config.GetInteger("maxSize")
    c.count = 0
    
    return c
}

func (c *MemoryCache) cleanup() {
    var oldest *CacheEntry
    now := time.Now()
    c.count = 0
    
    // Cleanup obsolete entries and find the oldest
    for key, value := range c.cache {
        diff := now.Sub(value.Created)
        // Remove obsolete entry
        if c.timeout > 0 && diff > c.timeout {
            delete(c.cache, key)
        } else {
            // Count the remaining entries
            c.count++
            if oldest != nil && oldest.Created.After(value.Created) {
                oldest = value
            }
        }
    }
    
    // Remove the oldest if cache size exceeded maximum
    if c.count > c.maxSize && oldest != nil {
        delete(c.cache, oldest.Key)
        c.count--
    }
}

func (c *MemoryCache) Get(key string, value interface{}) (interface{}, error) {
    entry := c.cache[key]
    
    // Cache has nothing
    if entry == nil { return nil, nil } 
    
    // Remove entry if expiration set and entry is expired
    diff := time.Now().Sub(entry.Created)
    if c.timeout > 0 && diff > c.timeout {
        delete(c.cache, key)
        c.count--
        return nil, nil
    }
    
    // Update access timeout
    entry.Accessed = time.Now()
    
    return entry.Value, nil
}

func (c *MemoryCache) Set(key string, value interface{}) (interface{}, error) {
    // Get the entry
    entry := c.cache[key]
    
    // Shortcut to remove entry from the cache
    if value == nil {
        if entry != nil {
            delete(c.cache, key)
            c.count--
        }
        return nil, nil
    }
    
    // Update the entry
    if entry != nil {
        now := time.Now()
        entry.Created = now
        entry.Accessed = now
        entry.Value = value
    } else {
        // Or create a new entry
        entry = NewCacheEntry(key, value)
        c.cache[key] = entry
        c.count++
    }
    
    // Clean up the cache
    if c.maxSize > 0 && c.count > c.maxSize {
        c.cleanup()
    }
    
    return value, nil
}

func (c *MemoryCache) Unset(key string) error {
    // Get the entry
    entry := c.cache[key]
    
    // Remove the entry from the cache
    if entry != nil {
        delete(c.cache, key)
        c.count --
    }
    
    return nil
}
