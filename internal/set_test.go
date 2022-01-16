package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	s1 := newSet()
	s1.add([]string{"a", "b", "c", "d"})

	s2 := newSet()
	s2.add([]string{"c"})

	s3 := newSet()
	s3.add([]string{"a", "c", "e"})

	assert.ElementsMatch(t, []string{"b", "d"}, s1.diff([]Set{s2, s3}))
}
