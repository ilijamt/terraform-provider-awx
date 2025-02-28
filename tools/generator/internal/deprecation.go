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
	&Deprecation2430{deprecation{name: "AssociateDisassociateGroups"}},
	&Deprecation2430{deprecation{name: "ObjectRoles"}},
}
