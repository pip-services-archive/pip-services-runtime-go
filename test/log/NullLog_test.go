package log

import (
    "testing"
    "github.com/stretchr/testify/suite"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/log"    
)

type NullLogTest struct {
    suite.Suite
    log runtime.ILog
    fixture *LogFixture
}

func (suite *NullLogTest) SetupTest() {
    suite.log = log.NewNullLog(nil)
    suite.fixture = NewLogFixture(suite.log)
}

func (suite *NullLogTest) TestLogLevel() {
    suite.fixture.TestLogLevel(suite.T())
}

func (suite *NullLogTest) TestTextOutput() {
    suite.fixture.TestTextOutput(suite.T())
}

func (suite *NullLogTest) TestMixedOutput() {
    suite.fixture.TestMixedOutput(suite.T())
}

func TestNullLogTestSuite(t *testing.T) {
    suite.Run(t, new(NullLogTest))
}