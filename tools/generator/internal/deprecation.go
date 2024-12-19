package internal

type deprecation struct {
	name string
}

func (d deprecation) Name() string {
	return d.name
}

type Deprecation interface {
	Name() string
	Check(mc *ModelConfig) (err error)
}

var deprecations = []Deprecation{
	&ObjectRole{deprecation{name: "ObjectRoles"}},
}
