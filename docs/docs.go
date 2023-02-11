package docs

import (
	_ "embed"
)

//go:embed openapi.swagger.json
var swaggerDocs string

type Docs struct{}

func (*Docs) ReadDoc() string { return swaggerDocs }
