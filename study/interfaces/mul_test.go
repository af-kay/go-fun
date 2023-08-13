package interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMul(t *testing.T) {
	assert.Equal(t, "HelloHelloHello", Mul("Hello", 3))
	assert.Equal(t, "Person LaraPerson Lara", Mul(person{"Lara"}, 2))
	assert.Equal(t, int64(40), Mul(int64(10), 4))
	assert.Equal(t, nil, Mul(nil, 9000))
}
