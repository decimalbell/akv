package internal

type Set map[string]struct{}

func newSet() Set {
	return make(Set)
}

func newSetWithSize(size int) Set {
	return make(Set, size)
}

func (s Set) add(members []string) int {
	count := 0
	for _, member := range members {
		if _, ok := s[member]; ok {
			continue
		}
		s[member] = struct{}{}
		count++
	}
	return count
}

func (s Set) remove(members []string) int {
	count := 0
	for _, member := range members {
		if _, ok := s[member]; ok {
			delete(s, member)
			count++
		}
	}
	return count
}

func (s Set) contains(member string) bool {
	_, ok := s[member]
	return ok
}

func any(ss []Set, fn func(s Set) bool) bool {
	for _, set := range ss {
		if fn(set) {
			return true
		}
	}
	return false
}

func (s Set) diff(ss []Set) []string {
	members := make([]string, 0, len(s))
	for member := range s {
		ok := any(ss, func(s Set) bool {
			return s.contains(member)
		})
		if !ok {
			members = append(members, member)
		}
	}
	return members
}

func (s Set) union(ss []Set) []string {
	size := s.len()
	for _, s := range ss {
		size += s.len()
	}

	result := newSetWithSize(size)
	for member := range s {
		result[member] = struct{}{}
	}
	for _, set := range ss {
		for member := range set {
			result[member] = struct{}{}
		}
	}

	return result.members()
}

func (s Set) members() []string {
	members := make([]string, 0, len(s))
	for member := range s {
		members = append(members, member)
	}
	return members
}

func (s Set) len() int {
	return len(s)
}
