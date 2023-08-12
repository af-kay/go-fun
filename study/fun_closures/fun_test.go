package funclosures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructClosures(t *testing.T) {
	s := makeMyStruct(20)

	assert.Equal(t, s.Log(""), "Struct value is 20, message is ", "Struct closure fn uses reference to it's owner")

	s.A = 40
	assert.Equal(t, s.Log(""), "Struct value is 40, message is ", "Struct closure fn sees owner updates")

	assert.Equal(t, s.Log("HELLO"), "Struct value is 40, message is HELLO", "Log message argument works")
}
