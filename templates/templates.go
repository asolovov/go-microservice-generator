package templates

import (
	"embed"
	"text/template"
)

//go:embed config/*.gotmpl
//go:embed internal/*.gotmpl
//go:embed internal/service/*.gotmpl
//go:embed internal/repository/*.gotmpl
//go:embed internal/server/*.gotmpl
//go:embed internal/server/grpcServer/*.gotmpl
//go:embed cmd/*.gotmpl
//go:embed cmd/root/*.gotmpl
//go:embed cmd/serve/*.gotmpl
var templatesFS embed.FS

var Tmpls *template.Template

func init() {
	var err error
	Tmpls, err = template.ParseFS(templatesFS,
		"config/*.gotmpl",
		"internal/*.gotmpl", "internal/service/*.gotmpl", "internal/repository/*.gotmpl",
		"internal/server/*.gotmpl", "internal/server/grpcServer/*.gotmpl",
		"cmd/root/*.gotmpl", "cmd/serve/*.gotmpl", "cmd/*.gotmpl",
	)
	if err != nil {
		panic(err)
	}
}
