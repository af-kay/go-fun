package datatypes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircularBuffer(t *testing.T) {
	buf := NewCircularBuffer(4)
	expected := [][]float64{nil, {1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}, {2, 3, 4, 5}}

	for i := 0; i < 6; i++ {
		if i > 0 {
			buf.AddValue(float64(i))
		}
		assert.EqualValues(t, expected[i], buf.GetValues())
	}
}
