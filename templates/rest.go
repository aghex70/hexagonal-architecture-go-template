package templates

const RestImportTemplate = `package server

import (
	"fmt"
	"{{.Module}}/config"
	"{{.Module}}/internal/core/ports"
`

const RestHandlerImportTemplate = `	"{{.Module}}/internal/handlers/{{.LowerEntity}}"
`

const RestServerTemplate = `)

type RestServer struct {
	cfg               config.RestConfig`

const RestServerHandlerTemplate = `
	{{.LowerEntity}}Handler {{.LowerEntity}}.{{.Entity}}Handler
	{{.LowerEntity}}Service ports.{{.Entity}}Servicer`

const StartServerTemplate = `
}

func (s *RestServer) StartServer() error {
	fmt.Println("Starting server")
	return nil
}

`

const NewRestServerTemplate = `func NewRestServer(cfg *config.RestConfig, `

const NewRestServerParamsTemplate = `{{.LowerEntity}}h {{.LowerEntity}}.{{.Entity}}Handler, `

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
