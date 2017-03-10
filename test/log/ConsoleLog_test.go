package log

import (
    "testing"
    "github.com/stretchr/testify/suite"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/log"    
)

type ConsoleLogTest struct {
    suite.Suite
    log runtime.ILog
    fixture *LogFixture
}

func (suite *ConsoleLogTest) SetupTest() {
    suite.log = log.NewConsoleLog(nil)
    suite.fixture = NewLogFixture(suite.log)
}

func (suite *ConsoleLogTest) TestLogLevel() {
    suite.fixture.TestLogLevel(suite.T())
}

func (suite *ConsoleLogTest) TestTextOutput() {
    suite.fixture.TestTextOutput(suite.T())
}

func (suite *ConsoleLogTest) TestMixedOutput() {
    suite.fixture.TestMixedOutput(suite.T())
}

func TestConsoleLogTestSuite(t *testing.T) {
    suite.Run(t, new(ConsoleLogTest))
}