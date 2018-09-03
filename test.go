package util

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func IsEqual(t *testing.T, got map[string]interface{}, field, want string)  {
	assert.Equal(t,got[field],want,"Expected "+ field +" to be '"+want+"'. Got "+got[field].(string))
}