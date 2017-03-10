package counters

import (
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/util"    
)

type LogCounters struct {
    AbstractCounters
}

func NewLogCounters(config *runtime.DynamicMap) *LogCounters {
    c := LogCounters { AbstractCounters: *NewAbstractCounters("LogCounters", config) }    
    c.SetSave(c.saveCounters)  
    return &c
}

func (c *LogCounters) counterToString(counter *runtime.Counter) string {
    output := "Counter " + counter.Name + " { "
    output += "\"type\": " + util.Converter.ToString(counter.Type)

    if counter.Last != nil {
        output += ", \"last\": " + util.Converter.ToString(*counter.Last)
    }
    if counter.Count != nil {
        output += ", \"count\": " + util.Converter.ToString(*counter.Count)
    }
    if counter.Min != nil {
        output += ", \"min\": " + util.Converter.ToString(*counter.Min)
    }
    if counter.Max != nil {
        output += ", \"max\": " + util.Converter.ToString(*counter.Max)
    }
    if counter.Avg != nil {
        output += ", \"avg\": " + util.Converter.ToString(*counter.Avg)
    }
    if counter.Time != nil {
        output += ", \"time\": " + util.Converter.ToString(*counter.Time)
    }

    output += " }"
    return output    
}

func (c *LogCounters) saveCounters(counters []*runtime.Counter) error {
    if len(counters) == 0 { return nil }
    
    for _, counter := range counters {
        c.Debug(c.counterToString(counter))
    }
    
    return nil
}