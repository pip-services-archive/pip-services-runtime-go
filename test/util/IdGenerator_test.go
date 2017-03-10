package util

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go/util"
)

func TestShort(t *testing.T) {
    id1 := util.IdGenerator.Short()
    assert.NotEmpty(t, id1)
    assert.True(t, len(id1) >= 9)
    
    id2 := util.IdGenerator.Short()
    assert.NotEmpty(t, id2)
    assert.True(t, len(id2) >= 9)
    
    assert.NotEqual(t, id1, id2)
}

func TestUUID(t *testing.T) {
    id1 := util.IdGenerator.UUID()
    assert.NotEmpty(t, id1)
    assert.Len(t, id1, 32)
    
    id2 := util.IdGenerator.UUID()
    assert.NotEmpty(t, id2)
    assert.Len(t, id2, 32)
    
    assert.NotEqual(t, id1, id2)
}