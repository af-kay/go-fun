package reflection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaiveIsNil(t *testing.T) {
	t.Run("Works good with nil & nil pointer", func(t *testing.T) {
		var actualNil *bool = nil

		// Why?
		assert.False(t, NaiveIsNil(actualNil))

		assert.True(t, NaiveIsNil(nil))
	})
	t.Run("Works bad with empty pointer", func(t *testing.T) {
		var emptyPtr *int

		assert.False(t, NaiveIsNil(emptyPtr))
	})
}

func TestIsNil(t *testing.T) {
	t.Run("Works good with nil & nil pointer", func(t *testing.T) {
		var actualNil *int = nil

		assert.True(t, IsNil(actualNil))
		assert.True(t, IsNil(nil))
	})
	t.Run("Works good with empty pointer", func(t *testing.T) {
		var emptyPtr *int

		assert.True(t, IsNil(emptyPtr))
	})
	t.Run("Works good with non-pointers and their addresses", func(t *testing.T) {
		var truthyValues = []interface{}{"String", 10.0, 42, true}

		for _, truthy := range truthyValues {
			assert.False(t, IsNil(truthy))
			assert.False(t, IsNil(&truthy))
		}
	})
	t.Run("Works good with falsy non-pointers and their addresses", func(t *testing.T) {
		var falsyNonNills = []interface{}{"", 0, 0.0, false}

		for _, falsy := range falsyNonNills {
			assert.False(t, IsNil(falsy))
			assert.False(t, IsNil(&falsy))
		}
	})
}
