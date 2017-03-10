package log

import (
    "testing"
    "github.com/stretchr/testify/suite"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/log"    
)

type CompositeLogTest struct {
    suite.Suite
    log runtime.ILog
    fixture *LogFixture
}

func (suite *CompositeLogTest) SetupTest() {
    logConfig1 := runtime.DynamicMap{ "level": 6 }
    log1 := log.NewConsoleLog(&logConfig1)
    log2 := log.NewNullLog(nil)
    
    logs := []runtime.ILog { log1, log2 }
    
    suite.log = log.NewCompositeLog(logs)
    suite.fixture = NewLogFixture(suite.log)
}

func (suite *CompositeLogTest) TestLogLevel() {
    suite.fixture.TestLogLevel(suite.T())
}

func (suite *CompositeLogTest) TestTextOutput() {
    suite.fixture.TestTextOutput(suite.T())
}

func (suite *CompositeLogTest) TestMixedOutput() {
    suite.fixture.TestMixedOutput(suite.T())
}

func TestCompositeLogTestSuite(t *testing.T) {
    suite.Run(t, new(CompositeLogTest))
}