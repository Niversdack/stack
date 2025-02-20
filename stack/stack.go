package stack

type Stack[t any] struct {
	a []t
}

func NewStack[t any](size int) Stack[t] {
	s := make([]t, size)
	return Stack[t]{
		a: s,
	}
}

func (s *Stack[t]) Pop() (r t, c bool) {
	if s.len() == 0 {
		return r, c
	}
	s.a, r = s.a[:s.len()-1], s.a[s.len()-1]
	return r, true
}

func (s *Stack[t]) Top() (r t, c bool) {
	if s.len() == 0 {
		return r, c
	}
	return s.a[s.len()-1], true
}

func (s *Stack[t]) len() int {
	return len(s.a)
}

func (s *Stack[t]) Push(T t) {
	s.a = append(s.a, T)
}

func (s *Stack[t]) GetStack() []t {
	return s.a
}
