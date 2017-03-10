// Todo: This implementation is not thread safe. Handle concurrency later

package log

import (
    "math"
    "time"
    "github.com/pip-services/pip-services-runtime-go"
)

type CachedLog struct {
    AbstractLog
    
    cache []LogEntry
    ticker *time.Ticker
}

type CachedLogSaver interface {
    Save(entries []LogEntry)
}

func NewCachedLog(id string, config *runtime.DynamicMap) *CachedLog {
    c := CachedLog{ AbstractLog: *NewAbstractLog(id, config) }
    c.cache = []LogEntry{}
    return &c
}

func (c *CachedLog) push(level int, correlationId string, message []interface{}) {
    if c.level < level { return }
    
    e := MakeLogEntry(c.Id(), level, correlationId, message)
    c.cache = append(c.cache, e)
}

func (c *CachedLog) popAll() []LogEntry {
    r := c.cache
    c.cache = []LogEntry{}
    return r
}

func (c *CachedLog) periodicSave() error {
    if len(c.cache) == 0 { return nil }
    
    entries := c.popAll()
    return c.Save(entries)
}

func (c *CachedLog) Save(entries []LogEntry) error {
    panic("Abstract method is called")
}

func (c *CachedLog) Open() error {
    timeout := time.Duration(math.Max(1000, c.Config().GetFloat("timeout")))
    
    // Stop previously started ticker
    if c.ticker != nil { c.ticker.Stop() }

    // Start a new ticker    
    c.ticker = time.NewTicker(time.Millisecond * timeout)
    
    go func() {
        for _ = range c.ticker.C {  
            err := c.periodicSave()
            if err != nil { c.Error(err) }
        }
    }()
    
    return c.AbstractComponent.Open()
}

func (c *CachedLog) Close() error {
    // Stop and clear ticker
    if c.ticker != nil {
        c.ticker.Stop()
        c.ticker = nil
    }
    
    err := c.periodicSave()
    if err != nil { return err }
    
    return c.AbstractComponent.Close()
}

func (c *CachedLog) Fatal(message ...interface{}) {
    c.push(runtime.LogLevelFatal, "", message)
}

func (c *CachedLog) Error(message ...interface{}) {
    c.push(runtime.LogLevelError, "", message)
}

func (c *CachedLog) Warn(message ...interface{}) {
    c.push(runtime.LogLevelWarn, "", message)
}

func (c *CachedLog) Info(message ...interface{}) {
    c.push(runtime.LogLevelInfo, "", message)
}

func (c *CachedLog) Debug(message ...interface{}) {
    c.push(runtime.LogLevelDebug, "", message)
}

func (c *CachedLog) Trace(message ...interface{}) {
    c.push(runtime.LogLevelTrace, "", message)
}
