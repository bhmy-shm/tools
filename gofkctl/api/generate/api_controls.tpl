package {{.pkgName}}

import (
	{{.imports}}
)

type {{.control}}Case struct {
	*wire.ServiceWire `inject:"-"`
}

func New{{.control}}Controller() *{{.control}}Case {
	return &{{.control}}Case{}
}

func (s *{{.control}}Case) Build(gofk *gofks.Gofk) {
    {{.handler}}
}

func (s *{{.control}}Case) Name() string {
	return "{{.control}}Case"
}

func (s *{{.control}}Case) Wire() *wire.ServiceWire {
	return s.ServiceWire
}
