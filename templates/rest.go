package templates

const RestImportTemplate = `package server

import (
	"{{.Module}}/config"
	"{{.Module}}/internal/core/ports"
`

const RestHandlerImportTemplate = `	"{{.Module}}/internal/handlers/{{.LowerEntity}}"
`

const RestServerTemplate = `	_ "{{.Module}}/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
)

// @title  API
// @version 1.0
// @description {{.ProjectName}} API server sample

// @contact.name API Support
// @contact.url https://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:12001
// @BasePath /api/v1
// @query.collection.format multi

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
	router.POST("/{{.TableSuffix}}", s.Handler.Create{{.Entity}})
	router.GET("/{{.TableSuffix}}/:uuid", s.Handler.Get{{.Entity}})
	router.PUT("/{{.TableSuffix}}/:uuid", s.Handler.Update{{.Entity}})
	router.DELETE("/{{.TableSuffix}}/:uuid", s.Handler.Delete{{.Entity}})
	router.GET("/{{.TableSuffix}}", s.Handler.List)

`

const NewRestServerTemplate = `	// swagger documentation
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return nil
}

func NewRestServer(cfg *config.RestConfig `

const NewRestServerParamsTemplate = `, {{.LowerEntity}}h {{.LowerEntity}}.Handler`

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
