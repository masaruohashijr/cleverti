package data

var exists = struct{}{}

type set struct {
	m map[interface{}]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[interface{}]struct{})
	return s
}

func (s *set) Add(value interface{}) {
	s.m[value] = exists
}

func (s *set) Remove(value interface{}) {
	delete(s.m, value)
}

func (s *set) Contains(value interface{}) bool {
	_, c := s.m[value]
	return c
}
