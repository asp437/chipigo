package main

type Stack struct {
	data []uint32
	i    uint16
	size uint16
}

func (s *Stack) init(size uint16) {
	s.data = make([]uint32, size)
	s.i = 0
	s.size = size
}

func (s *Stack) push(val uint32) {
	if s.i == s.size {
		return
	}
	s.data[s.i] = val
	s.i++
}

func (s *Stack) pop() uint32 {
	s.i--
	ret := s.data[s.i]
	return ret
}
