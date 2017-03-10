package counters

import (
    "testing"
    "github.com/stretchr/testify/suite"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/log"    
    "github.com/pip-services/pip-services-runtime-go/counters"    
)

type NullCountersTest struct {
    suite.Suite
    counters runtime.ICounters
}

func (suite *NullCountersTest) SetupTest() {
    clog := log.NewConsoleLog(nil)
    refs := runtime.NewReferences().WithLog(clog)
    
    suite.counters = counters.NewNullCounters(nil)
    suite.counters.Init(refs)
    suite.counters.Open()
}

func (suite *NullCountersTest) TearDownTest() {
    suite.counters.Close()
}

func (suite *NullCountersTest) TestSimpleCounters() {
    suite.counters.Last("Test.LastValue", 123)
    suite.counters.Increment("Test.Increment", 3)
    suite.counters.Stats("Test.Statistics", 123)
    
    suite.counters.Dump()
}

func (suite *NullCountersTest) TestMeasureElapsedTime() {
    timing := suite.counters.BeginTiming("Test.Elapsed")
    timing.EndTiming()
    
    suite.counters.Dump()
}

func TestNullCountersTestSuite(t *testing.T) {
    suite.Run(t, new(NullCountersTest))
}