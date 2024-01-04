package tests

import (
	"testing"

	"github.com/codegshinde/raptor"

	"github.com/stretchr/testify/assert"
)

func TestRaptor(t *testing.T) {
	// Test SetConfigFile and ReadConfigFile
	err := raptor.SetConfigFile("D:/go/probys-microservice/upload/.raptor")
	assert.NoError(t, err, "Unexpected error setting config file")

	// Test Get
	value, err := raptor.Get("BUCKET_NAME")
	assert.NoError(t, err, "Unexpected error getting value")
	assert.Equal(t, "swayamkendra-files", value, "Unexpected value for key 'yourKey'")

	// Test Set
	raptor.Set("anotherKey", "anotherValue")
	value, err = raptor.Get("anotherKey")
	assert.NoError(t, err, "Unexpected error getting value after set")
	assert.Equal(t, "anotherValue", value, "Unexpected value for key 'anotherKey' after set")

	// Test Delete
	raptor.Delete("yourKey")
	_, err = raptor.Get("yourKey")
	assert.Error(t, err, "Expected error for deleted key 'yourKey'")
	assert.EqualError(t, err, "key 'yourKey' not found in configuration", "Unexpected error message")

	// Test GetAll

	allValues := raptor.GetAll()
	assert.Len(t, allValues, 3, "Unexpected number of keys in configuration map")
	assert.Contains(t, allValues, "anotherKey", "Expected key 'anotherKey' in configuration map")
	assert.Equal(t, "anotherValue", allValues["anotherKey"], "Unexpected value for key 'anotherKey' in configuration map")
}
