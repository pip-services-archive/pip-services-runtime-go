package build

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/build"    
    "github.com/pip-services/pip-services-runtime-go/counters"    
    "github.com/pip-services/pip-services-runtime-go/log"    
)

type LifeCycleManagerTest struct {
    suite.Suite
    refs *runtime.References
}

func (suite *LifeCycleManagerTest) SetupTest() {
    nlog := log.NewNullLog(nil)
    ncounters := counters.NewNullCounters(nil)

    suite.refs = runtime.NewReferences().WithLog(nlog).WithCounters(ncounters)
}

func (suite *LifeCycleManagerTest) TestInit() {
    build.LifeCycleManager.Init(suite.refs)
}

func (suite *LifeCycleManagerTest) TestInitAndOpen() {
    err := build.LifeCycleManager.InitAndOpen(suite.refs)
    assert.Nil(suite.T(), err)
}

func (suite *LifeCycleManagerTest) TestOpen() {
    build.LifeCycleManager.Init(suite.refs)
    err := build.LifeCycleManager.Open(suite.refs)
    assert.Nil(suite.T(), err)
}

func (suite *LifeCycleManagerTest) TestClose() {
    err := build.LifeCycleManager.InitAndOpen(suite.refs)
    assert.Nil(suite.T(), err)

    err = build.LifeCycleManager.Close(suite.refs)
    assert.Nil(suite.T(), err)
}

func (suite *LifeCycleManagerTest) TestForceClose() {
    err := build.LifeCycleManager.InitAndOpen(suite.refs)
    assert.Nil(suite.T(), err)

    err = build.LifeCycleManager.Close(suite.refs)
    assert.Nil(suite.T(), err)
}

func TestLifeCycleManagerTestSuite(t *testing.T) {
    suite.Run(t, new(LifeCycleManagerTest))
}