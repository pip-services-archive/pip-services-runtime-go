package api

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/build"    
    "github.com/pip-services/pip-services-runtime-go/test/db"    
    "github.com/pip-services/pip-services-runtime-go/test/logic"    
)

type DummyRestServiceTest struct {
    suite.Suite
    
    db *db.DummyFileDataAccess
    ctrl *logic.DummyController
    api *DummyRestService
    refs *runtime.References
}

func (suite *DummyRestServiceTest) SetupSuite() {
    dbConfig := runtime.NewMapAndSet(
        "type", "file",
        "path", "../../dummies.json",
        "data", []*db.Dummy {}, 
    )    
    suite.db = db.NewDummyFileDataAccess(dbConfig)    
    suite.ctrl = logic.NewDummyController(nil)

    apiConfig := runtime.NewMapAndSet(
        "type", "rest",
        "transport.type", "http",
        "transport.host", "localhost",
        "transport.port", 3000,
    )
    suite.api = NewDummyRestService(apiConfig)
    
    suite.refs = runtime.NewReferences().WithDB(suite.db).WithCtrl(suite.ctrl).WithAPI(suite.api)
    
    err := build.LifeCycleManager.InitAndOpen(suite.refs)
    assert.Nil(suite.T(), err)
}

func (suite *DummyRestServiceTest) SetupTest() {
    suite.db.Clear()
}

func (suite *DummyRestServiceTest) TearDownSuite() {
    err := build.LifeCycleManager.Close(suite.refs)
    assert.Nil(suite.T(), err)
}

func (suite *DummyRestServiceTest) TestCrudOperations() {
    assert.True(suite.T(), true)
}

func TestDummyRestServiceTestSuite(t *testing.T) {
    suite.Run(t, new(DummyRestServiceTest))
}