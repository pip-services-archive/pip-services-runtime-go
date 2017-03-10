package counters

import (
    "testing"
    "github.com/stretchr/testify/suite"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/log"    
    "github.com/pip-services/pip-services-runtime-go/counters"    
)

type LogCountersTest struct {
    suite.Suite
    counters runtime.ICounters
    fixture *CountersFixture
}

func (suite *LogCountersTest) SetupTest() {
    clog := log.NewConsoleLog(nil)
    refs := runtime.NewReferences().WithLog(clog)
    
    suite.counters = counters.NewLogCounters(nil)
    suite.counters.Init(refs)
    suite.counters.Open()
    
    suite.fixture = NewCountersFixture(suite.counters)
}

func (suite *LogCountersTest) TearDownTest() {
    suite.counters.Close()
}

func (suite *LogCountersTest) TestSimpleCounters() {
    suite.fixture.TestSimpleCounters(suite.T())
}

// func (suite *LogCountersTest) TestMeasureElapsedTime() {
//     suite.fixture.TestMeasureElapsedTime(suite.T())    
// }

func TestLogCountersTestSuite(t *testing.T) {
    suite.Run(t, new(LogCountersTest))
}