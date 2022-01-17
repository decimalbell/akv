package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	s := newSet()
	s.add([]string{"a", "a", "b"})

	assert.ElementsMatch(t, []string{"a", "b"}, s.members())
}

func TestRemove(t *testing.T) {
	s := newSet()
	s.add([]string{"a", "b"})
	s.remove([]string{"a"})

	assert.ElementsMatch(t, []string{"b"}, s.members())
}

func TestDiff(t *testing.T) {
	s1 := newSet()
	s1.add([]string{"a", "b", "c", "d"})

	s2 := newSet()
	s2.add([]string{"c"})

	s3 := newSet()
	s3.add([]string{"a", "c", "e"})

	assert.ElementsMatch(t, []string{"b", "d"}, s1.diff([]Set{s2, s3}))
}

func TestUnion(t *testing.T) {
	s1 := newSet()
	s1.add([]string{"a", "b", "c", "d"})

	s2 := newSet()
	s2.add([]string{"c"})

	s3 := newSet()
	s3.add([]string{"a", "c", "e"})

	assert.ElementsMatch(t, []string{"a", "b", "c", "d", "e"}, s1.union([]Set{s2, s3}))
}
