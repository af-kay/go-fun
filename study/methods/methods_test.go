package methods

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAgeString(t *testing.T) {
    var a Age = 20

    assert.Equal(t, a.String(), "Age: 20")
}
