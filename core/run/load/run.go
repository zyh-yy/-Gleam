package registry

func (s *Loader) Do() {
	s.Input = make([]interface{}, s.InputLen)
	s.Output = make([]interface{}, s.OutputLen)
	for i := 0; i < len(s.Input); i++ {
		data := <-s.Node.Input[i]
		s.Input = append(s.Input, data)
	}

	// 处理逻辑
	err := s.fun()
	if err != nil {
		return
	}

	for i := 0; i < len(s.Input); i++ {
		s.Node.Output[i] <- s.Output[i]
	}
}
