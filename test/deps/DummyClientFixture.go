package deps

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-services-runtime-go/test/db"
)

type DummyClientFixture struct {
    client IDummyClient
}

func NewDummyClientFixture(client IDummyClient) *DummyClientFixture {
    return &DummyClientFixture { client: client }    
}

var DUMMY1 *db.Dummy = db.NewDummy("", "Key 1", "Content 1")
var DUMMY2 *db.Dummy = db.NewDummy("", "Key 2", "Content 2")

func (c *DummyClientFixture) TestCrudOperations(t *testing.T) {
    // Create one dummy
    dummy1, err := c.client.CreateDummy(DUMMY1)    
    assert.Nil(t, err)
    
    assert.NotNil(t, dummy1)
    assert.NotEmpty(t, dummy1.ID)
    assert.Equal(t, DUMMY1.Key, dummy1.Key)
    assert.Equal(t, DUMMY1.Content, dummy1.Content)

    // Create another dummy
    dummy2, err := c.client.CreateDummy(DUMMY2)    
    assert.Nil(t, err)
    
    assert.NotNil(t, dummy2)
    assert.NotEmpty(t, dummy2.ID)
    assert.Equal(t, DUMMY2.Key, dummy2.Key)
    assert.Equal(t, DUMMY2.Content, dummy2.Content)
    
    // Get all dummies
    dummies, err2 := c.client.GetDummies(nil, nil)
    assert.Nil(t, err2)
    assert.NotNil(t, dummies)
    assert.NotNil(t, dummies.Data)
    assert.Len(t, dummies.Data, 2)
        
    // Update the dummy
    dummyData := runtime.NewMapAndSet("content", "Updated Content 1")
    dummy1, err = c.client.UpdateDummy(dummy1.ID, dummyData)
    assert.Nil(t, err)
    assert.NotNil(t, dummy1)
    assert.Equal(t, "Updated Content 1", dummy1.Content)

    // Delete the dummy
    err = c.client.DeleteDummy(dummy1.ID)
    assert.Nil(t, err)
    
    // Try to get deleted dummy
    dummy1, err = c.client.GetDummyById(dummy1.ID)
    assert.Nil(t, err)
    assert.Nil(t, dummy1) 
}