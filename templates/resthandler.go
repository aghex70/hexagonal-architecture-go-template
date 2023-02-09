package templates

const RestHandlerTemplate = `package {{.LowerEntity}}

import (
    "github.com/gin-gonic/gin"
	"{{.Module}}/internal/core/ports"
    "net/http"
)

type {{.Entity}}Handler struct {
	{{.LowerEntity}}Service ports.{{.Entity}}Servicer
}

func (h {{.Entity}}Handler) Create{{.Entity}}(c *gin.Context) {
	n{{.Initial}}, err := h.{{.LowerEntity}}Service.Create(nil, r, payload)
	if err != nil {
		panic(err)
	}
}

func (h {{.Entity}}Handler) Update{{.Entity}}(c *gin.Context) {
	err := h.{{.LowerEntity}}Service.Update(nil, r, payload)
	if err != nil {
		panic(err)
	}
}

func (h {{.Entity}}Handler) Get{{.Entity}}(c *gin.Context) {
	{{.Initial}}{{.Initial}}, err := h.{{.LowerEntity}}Service.Get(nil, r, payload)
	if err != nil {
		panic(err)
	}
}

func (h {{.Entity}}Handler) Delete{{.Entity}}(c *gin.Context) {
	err := h.{{.LowerEntity}}Service.Delete(nil, r, payload)
	if err != nil {
		panic(err)
	}
}

func (h {{.Entity}}Handler) List(c *gin.Context) {
	{{.Initial}}s, err := h.{{.LowerEntity}}Service.List(nil, r)
	if err != nil {
		panic(err)
	}
}

func New{{.Entity}}Handler(cs ports.{{.Entity}}Servicer) {{.Entity}}Handler {
	return {{.Entity}}Handler{
		{{.LowerEntity}}Service: {{.Initial}}s,
	}
}
`

func GetRestHandlerFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        RestHandlerTemplate,
			TemplateContext: tc,
		},
	}
}
