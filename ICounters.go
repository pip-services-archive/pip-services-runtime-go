package runtime

import "time"

const (
    CounterTypeInterval = 0
    CounterTypeLastValue = 1
    CounterTypeStatistics = 2
    CounterTypeTimestamp = 3
    CounterTypeIncrement = 4
)

type ICounters interface{
    IComponent
    
    Reset(name string)
    ResetAll()
    Dump() error
    
    GetAll() []*Counter
    Get(name string, typ int) *Counter
    
    BeginTiming(name string) *Timing
    Stats(name string, value float32)
    Last(name string, value float32)
    TimestampNow(name string)
    Timestamp(name string, value time.Time)
    IncrementOne(name string)
    Increment(name string, value int)
}