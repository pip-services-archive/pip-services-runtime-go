package cache

import "time"

type CacheEntry struct {
    Created         time.Time
    Accessed        time.Time
    Key             string
    Value           interface{}
}

func NewCacheEntry(key string, value interface{}) *CacheEntry {
    now := time.Now();
    
    return &CacheEntry {
        Created: now,
        Accessed: now,
        Key: key,
        Value: value,
    }
}