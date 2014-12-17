package condorcet

type sorter struct {
	lenFunc  func() int
	lessFunc func(i, j int) bool
	swapFunc func(i, j int)
}

func (s *sorter) Len() int {
	return s.lenFunc()
}

func (s *sorter) Less(i, j int) bool {
	return s.lessFunc(i, j)
}

func (s *sorter) Swap(i, j int) {
	s.swapFunc(i, j)
}
