package encryption

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyGeneration(t *testing.T) {
	key1, _ := NewKey()
	key2, _ := NewKey()
	assert.Equal(t, len(key1), KeySize)
	assert.Equal(t, len(key2), KeySize)
	assert.NotEqual(t, key1, key2)
}
