package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	val := []byte("bytes")
	value := newBytesValue(val)

	bytes, err := value.bytes()
	assert.Nil(t, err)
	assert.EqualValues(t, val, bytes)

	_, err = value.set()
	assert.NotNil(t, err)
}

func TestHash(t *testing.T) {
	value := newHashValue()

	_, err := value.hash()
	assert.Nil(t, err)

	_, err = value.bytes()
	assert.NotNil(t, err)
}

func TestSet(t *testing.T) {
	value := newSetValue()

	_, err := value.set()
	assert.Nil(t, err)

	_, err = value.hash()
	assert.NotNil(t, err)
}
