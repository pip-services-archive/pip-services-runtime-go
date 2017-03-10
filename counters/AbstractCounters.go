// Todo: This implementation is not thread safe. Handle concurrency later

package counters

import (
    "math"
    "time"
    "github.com/pip-services/pip-services-runtime-go"
)

/********** AbstractCounters class ********/

type AbstractCounters struct {
    runtime.AbstractComponent

    cache map[string]*runtime.Counter
    updated bool
    ticker *time.Ticker  
        
    // This is a hack to work around lack of abstract methods
    save func(counters []*runtime.Counter) error  
}

func NewAbstractCounters(id string, config *runtime.DynamicMap) *AbstractCounters {
    defaultConfig := runtime.DynamicMap { "timeout": 60000 }
    config = runtime.NewMapWithDefaults(config, &defaultConfig)
    
    c := AbstractCounters { AbstractComponent: *runtime.NewAbstractComponent(id, config) }
    c.cache = map[string]*runtime.Counter {}    
    c.save = c.saveCounters      
    return &c
}

func (c *AbstractCounters) SetSave(save func(counters []*runtime.Counter) error) {
    c.save = save
} 

func (c *AbstractCounters) Open() error {
    timeout := time.Duration(math.Max(1000, c.Config().GetFloat("timeout")))
    
    // Stop previously started ticker
    if c.ticker != nil { c.ticker.Stop() }

    // Start a new ticker    
    ticker := time.NewTicker(time.Millisecond * timeout)
    c.ticker = ticker
    
    go func() {
        for _ = range ticker.C {  
            err := c.Dump()
            if err != nil { c.Error(err) }
        }
    }()
    
    return c.AbstractComponent.Open()
}

func (c *AbstractCounters) Close() error {
    // Stop and clear ticker
    if c.ticker != nil {
        c.ticker.Stop()
        c.ticker = nil
    }
    
    // Save and clear counters if any
    if c.updated {
        counters := c.GetAll()
        c.ResetAll()

        err := c.save(counters)
        if err != nil { return err }
    }
    
    return c.AbstractComponent.Close()
}

func (c *AbstractCounters) saveCounters(counters []*runtime.Counter) error {
    panic("Called abstract method")
}

func (c *AbstractCounters) Reset(name string) {
    delete(c.cache, name)
}

func (c *AbstractCounters) ResetAll() {
    c.cache = map[string]*runtime.Counter {}
    c.updated = false
}

func (c *AbstractCounters) Dump() error {
    if c.updated {
        counters := c.GetAll()
        return c.save(counters)
    }
    return nil
}

func (c *AbstractCounters) GetAll() []*runtime.Counter {
    counters := make([]*runtime.Counter, 0, len(c.cache))    
    for _, c := range c.cache {
        counters = append(counters, c)
    }
    return counters
}

func (c *AbstractCounters) Get(name string, typ int) *runtime.Counter {
    if name == "" { panic("Counter name was not set") }
    
    counter, ok := c.cache[name]
    
    if ok == false || counter.Type != typ {
        counter = runtime.NewCounter(name, typ)
        c.cache[name] = counter
    }
    
    return counter
}

func (c *AbstractCounters) calculateStats(counter *runtime.Counter, value float32) {
    if counter == nil { panic("Missing counter") }

    last := value
    counter.Last = &last
    
    if counter.Count != nil { 
        *counter.Count = *counter.Count + 1  
    } else {
        count := 1
        counter.Count = &count
    }
    
    if counter.Max != nil {
        *counter.Max = float32(math.Max(float64(*counter.Max), float64(value)))
    } else {
        max := value
        counter.Max = &max
    }
    
    if counter.Min != nil {
        *counter.Max = float32(math.Min(float64(*counter.Min), float64(value)))
    } else {
        min := value
        counter.Min = &min
    }
    
    if counter.Avg != nil && *counter.Count > 1 {
        *counter.Avg = ((*counter.Avg) * (float32(*counter.Count) - 1) + value) / float32(*counter.Count) 
    } else {
        avg := value
        counter.Avg = &avg
    }
    
    c.updated = true
}

func (c *AbstractCounters) SetTiming(name string, elapsed float32) {
    counter := c.Get(name, runtime.CounterTypeInterval)
    c.calculateStats(counter, elapsed)
}

func (c *AbstractCounters) BeginTiming(name string) *runtime.Timing {
    return runtime.NewTiming(c, name)
}

func (c *AbstractCounters) Stats(name string, value float32) {
    counter := c.Get(name, runtime.CounterTypeStatistics)
    c.calculateStats(counter, value)
}

func (c *AbstractCounters) Last(name string, value float32) {
    counter := c.Get(name, runtime.CounterTypeLastValue)
    counter.Last = &value
    c.updated = true
}

func (c *AbstractCounters) TimestampNow(name string) {
    c.Timestamp(name, time.Now())
}

func (c *AbstractCounters) Timestamp(name string, time time.Time) {
    counter := c.Get(name, runtime.CounterTypeTimestamp)
    time = time.UTC()
    counter.Time = &time
    c.updated = true
}

func (c *AbstractCounters) IncrementOne(name string) {
    c.Increment(name, 1)
}

func (c *AbstractCounters) Increment(name string, value int) {
    counter := c.Get(name, runtime.CounterTypeIncrement)
    if counter.Count != nil {
        *counter.Count = *counter.Count + value
    } else {
        count := 1
        counter.Count = &count
    }
    c.updated = true
}
