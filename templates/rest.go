package templates

const RestImportTemplate = `package server

import (
	"{{.Module}}/config"
	"{{.Module}}/internal/core/ports"
`

const RestHandlerImportTemplate = `	"{{.Module}}/internal/handlers/{{.LowerEntity}}"
`

const RestServerTemplate = `	"github.com/gin-gonic/gin"
)

type RestServer struct {
	cfg               config.RestConfig`

const RestServerHandlerTemplate = `
	{{.LowerEntity}}Handler {{.LowerEntity}}.{{.Entity}}Handler
	{{.LowerEntity}}Service ports.{{.Entity}}Servicer`

const StartServerTemplate = `
}

func (s *RestServer) StartServer() error {
	router := gin.Default()

`

const StartServerHandlersTemplate = `	// {{.TableSuffix}}
	router.POST("/{{.TableSuffix}}", s.{{.LowerEntity}}Handler.Create{{.Entity}})
	router.GET("/{{.TableSuffix}}/:uuid", s.{{.LowerEntity}}Handler.Get{{.Entity}})
	router.PUT("/{{.TableSuffix}}/:uuid", s.{{.LowerEntity}}Handler.Update{{.Entity}})
	router.DELETE("/{{.TableSuffix}}/:uuid", s.{{.LowerEntity}}Handler.Delete{{.Entity}})
	router.GET("/{{.TableSuffix}}", s.{{.LowerEntity}}Handler.List)

`

const NewRestServerTemplate = `	return nil
}

func NewRestServer(cfg *config.RestConfig `

const NewRestServerParamsTemplate = `, {{.LowerEntity}}h {{.LowerEntity}}.{{.Entity}}Handler`

const NewRestServerParamsContinueTemplate = `) *RestServer {
	return &RestServer{
		cfg:               *cfg,`

const NewRestServerHandlerTemplate = `
		{{.LowerEntity}}Handler: {{.LowerEntity}}h,`

const NewRestServerEndTemplate = `
		}
}`

func GetRestServerFileConfiguration(entities []string, tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        RestImportTemplate,
			TemplateContext: tc,
		},
		{
			Template:        RestHandlerImportTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
		{
			Template:        RestServerTemplate,
			TemplateContext: tc,
		},
		{
			Template:        RestServerHandlerTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
		{
			Template:        StartServerTemplate,
			TemplateContext: tc,
		},
		{
			Template:        StartServerHandlersTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
		{
			Template:        NewRestServerTemplate,
			TemplateContext: tc,
		},
		{
			Template:        NewRestServerParamsTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
		{
			Template:        NewRestServerParamsContinueTemplate,
			TemplateContext: tc,
		},
		{
			Template:        NewRestServerHandlerTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
		{
			Template:        NewRestServerEndTemplate,
			TemplateContext: tc,
		},
	}
}
