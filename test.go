package util

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func IsEqual(t *testing.T, want, field string ,got interface{})  {
	assert.Equal(t,got,want,"Expected "+ field +" to be '"+want+"'. Got "+got.(string))
}