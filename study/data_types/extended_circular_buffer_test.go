package datatypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtendedCircularBuffer(t *testing.T) {
	buffer := NewExtendedCircularBuffer(3)

	buffer.addValues(1, 2, 3, 4, 5)

	assert.EqualValues(t, buffer.GetValues(), []float64{3, 4, 5})
}
