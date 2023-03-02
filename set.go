package set

type Set map[string]bool

func NewSet() Set {
	s := make(Set)
	return s
}

func (s Set) Add(item string) {
	s[item] = true
}

func (s Set) Remove(item string) {
	delete(s, item)
}

func (s Set) Contains(item string) bool {
	_, ok := s[item]
	return ok
}

func (s Set) Size() int {
	return len(s)
}
