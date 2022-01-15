package internal

type Map map[string][]byte

func newMap() Map {
	return make(Map)
}

func (m Map) len() int {
	return len(m)
}

func (m Map) keys() []string {
	strs := make([]string, 0, len(m))
	for key := range m {
		strs = append(strs, key)
	}
	return strs
}

func (m Map) values() [][]byte {
	strs := make([][]byte, 0, len(m))
	for _, value := range m {
		strs = append(strs, value)
	}
	return strs
}

func (m Map) items() [][]byte {
	strs := make([][]byte, 0, 2*len(m))
	for key, value := range m {
		strs = append(strs, []byte(key), value)
	}
	return strs
}
