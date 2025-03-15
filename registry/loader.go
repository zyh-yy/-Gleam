package registry

import (
	"runflex/core"
)

type Loader struct {
	core.Node
	Input  []interface{}
	Output []interface{}
}

func (s *Loader) Do() {
	s.Input = make([]interface{}, s.InputLen)
	s.Output = make([]interface{}, s.OutputLen)
	for i := 0; i < len(s.Input); i++ {
		data := <-s.Node.Input[i]
		s.Input = append(s.Input, data)
	}

	// 处理逻辑

	for i := 0; i < len(s.Input); i++ {
		s.Node.Output[i] <- s.Output[i]
	}
}
