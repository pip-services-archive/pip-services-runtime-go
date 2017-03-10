package log

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type CompositeLog struct {
    AbstractLog
    
    logs []runtime.ILog
}

func NewCompositeLog(logs []runtime.ILog) *CompositeLog {
    c := CompositeLog{ AbstractLog: *NewAbstractLog("CompositeLog", nil) }
    c.logs = logs
    c.level = c.maxLevel(logs)
    return &c
}

func (c *CompositeLog) maxLevel(logs []runtime.ILog) int {
    if len(logs) == 0 { return runtime.LogLevelNone }
    
    level := runtime.LogLevelNone
    for _, log := range logs {
        if log.Level() > level { 
            level = log.Level()    
        }
    }
    
    return level
}

func (c *CompositeLog) Init(refs *runtime.References) error {
    // Initialize subloggers
    for _, log := range c.logs {
        if err := log.Init(refs); err != nil { return err } 
    }
    
    c.level = c.maxLevel(c.logs)
    return c.AbstractLog.Init(refs)    
}

func (c *CompositeLog) Open() error {
    // Open subloggers
    for _, log := range c.logs {
        if err := log.Open(); err != nil { return err } 
    }
    
    return c.AbstractLog.Open()    
}

func (c *CompositeLog) Close() error {
    // Close subloggers
    for _, log := range c.logs {
        if err := log.Close(); err != nil { return err } 
    }
    
    return c.AbstractLog.Close()    
}

func (c *CompositeLog) Fatal(message ...interface{}) {
    if (c.level >= runtime.LogLevelFatal) {
        for _, log := range c.logs {
            log.Fatal(message)
        }
    }
}

func (c *CompositeLog) Error(message ...interface{}) {
    if (c.level >= runtime.LogLevelError) {
        for _, log := range c.logs {
            log.Error(message)
        }
    }
}

func (c *CompositeLog) Warn(message ...interface{}) {
    if (c.level >= runtime.LogLevelWarn) {
        for _, log := range c.logs {
            log.Warn(message)
        }
    }
}

func (c *CompositeLog) Info(message ...interface{}) {
    if (c.level >= runtime.LogLevelInfo) {
        for _, log := range c.logs {
            log.Info(message)
        }
    }
}

func (c *CompositeLog) Debug(message ...interface{}) {
    if (c.level >= runtime.LogLevelDebug) {
        for _, log := range c.logs {
            log.Debug(message)
        }
    }
}

func (c *CompositeLog) Trace(message ...interface{}) {
    if (c.level >= runtime.LogLevelTrace) {
        for _, log := range c.logs {
            log.Trace(message)
        }
    }
}
