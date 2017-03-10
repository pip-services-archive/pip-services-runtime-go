package counters

import (
    "time"
    "github.com/pip-services/pip-services-runtime-go"    
)

type NullCounters struct {
    runtime.AbstractComponent
}

func NewNullCounters(config *runtime.DynamicMap) *NullCounters {
    return &NullCounters { AbstractComponent: *runtime.NewAbstractComponent("NullCounters", config) }
}

func (c *NullCounters) Reset(name string) {}
func (c *NullCounters) ResetAll() {}

func (c *NullCounters) Dump() error { 
    return nil 
}

func (c *NullCounters) GetAll() []*runtime.Counter { 
    return []*runtime.Counter {} 
}

func (c *NullCounters) Get(name string, typ int) *runtime.Counter {
    if name == "" { panic("Counter name was not set") }    
    return runtime.NewCounter(name, typ)
}

func (c *NullCounters) BeginTiming(name string) *runtime.Timing {
    return runtime.EmptyTiming()
}

func (c *NullCounters) Stats(name string, value float32) {}
func (c *NullCounters) Last(name string, value float32) {}
func (c *NullCounters) TimestampNow(name string) {}
func (c *NullCounters) Timestamp(name string, time time.Time) {}
func (c *NullCounters) IncrementOne(name string) {}
func (c *NullCounters) Increment(name string, value int) {}
