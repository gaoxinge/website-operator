package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt32Ptr(t *testing.T) {
	intPtr := Int32Ptr(1)
	assert.Equal(t, int32(1), *intPtr)
}
