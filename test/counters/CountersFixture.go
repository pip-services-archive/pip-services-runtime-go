package counters

import (
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go"    
)

type CountersFixture struct {
    counters runtime.ICounters
}

func NewCountersFixture(counters runtime.ICounters) *CountersFixture {
    return &CountersFixture{ counters: counters }
}

func (c *CountersFixture) TestSimpleCounters(t *testing.T) {
    c.counters.Last("Test.LastValue", 123)
    c.counters.Last("Test.LastValue", 123456)

    counter := c.counters.Get("Test.LastValue", runtime.CounterTypeLastValue)
    assert.NotNil(t, counter)
    assert.Equal(t, float32(123456.), *counter.Last)

    c.counters.IncrementOne("Test.Increment")
    c.counters.Increment("Test.Increment", 3)

    counter = c.counters.Get("Test.Increment", runtime.CounterTypeIncrement)
    assert.NotNil(t, counter.Count)
    assert.Equal(t, int(4), *counter.Count)

    c.counters.TimestampNow("Test.Timestamp")
    c.counters.TimestampNow("Test.Timestamp")

    counter = c.counters.Get("Test.Timestamp", runtime.CounterTypeTimestamp)
    assert.NotNil(t, counter)
    assert.NotNil(t, counter.Time)

    c.counters.Stats("Test.Statistics", 1)
    c.counters.Stats("Test.Statistics", 2)
    c.counters.Stats("Test.Statistics", 3)

    counter = c.counters.Get("Test.Statistics", runtime.CounterTypeStatistics)
    assert.NotNil(t, counter.Avg)
    assert.Equal(t, float32(2.), *counter.Avg)

    c.counters.Dump()
}

func (c *CountersFixture) TestMeasureElapsedTime(t *testing.T) {
    timing := c.counters.BeginTiming("Test.Elapsed")
    time.Sleep(time.Millisecond * 100)
    timing.EndTiming()
    
    counter := c.counters.Get("Test.Elapsed", runtime.CounterTypeInterval)
    assert.NotNil(t, counter)
    assert.True(t, *counter.Last > 50)
    
    assert.True(t, *counter.Last < 5000)
    
    c.counters.Dump()
}
