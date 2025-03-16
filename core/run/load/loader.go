package registry

import (
	"runflex/core"
)

type LoadFunc func() error

type Loader struct {
	core.Node
	Input  []interface{}
	Output []interface{}
	fun    func() error
}
