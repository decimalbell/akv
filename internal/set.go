package internal

type Set map[string]struct{}

func newSet() Set {
	return make(Set)
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

func (s Set) diff(ss []Set) []string {
	members := make([]string, 0, len(s))
	for member := range s {
		ok := false
		for _, set := range ss {
			if set.contains(member) {
				ok = true
				break
			}
		}
		if !ok {
			members = append(members, member)
		}
	}
	return members
}

func (s Set) len() int {
	return len(s)
}
