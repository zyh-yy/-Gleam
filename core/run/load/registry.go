package registry

func (l *Loader) NewLoader(fun LoadFunc) *Loader {
	return &Loader{
		fun: fun,
	}
}

func (l *Loader) WithDep(key string) {
}
