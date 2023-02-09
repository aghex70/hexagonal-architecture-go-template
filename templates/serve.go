package templates

const ServeStartTemplate = `package cmd

import (
	"{{.Module}}/{{.ProjectName}}/config"
	"{{.Module}}/{{.ProjectName}}/persistence/database"
	"github.com/spf13/cobra"
	"log"
`

const ServeImportTemplate = `	{{.LowerEntity}}Service "{{.Module}}/{{.ProjectName}}/internal/core/services/{{.LowerEntity}}"
	{{.LowerEntity}}Handler "{{.Module}}/{{.ProjectName}}/internal/handlers/{{.LowerEntity}}"
	{{.LowerEntity}}Repository "{{.Module}}/{{.ProjectName}}/internal/repositories/{{.LowerEntity}}"
`

const ServeCommandTemplate = `)

func ServeCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve application",
		Run: func(cmd *cobra.Command, args []string) {
			gdb, err := database.NewGormDB(*cfg.Database)
			if err != nil {
				panic(err)
			}
`

const ServeRepositoryTemplate = `			{{.Initial}}r := {{.LowerEntity}}Repository.New{{.Entity}}GormRepository(gdb)
`

const ServeServiceAndHandlerTemplate = `
			{{.Initial}}s := {{.LowerEntity}}Service.New{{.Entity}}Service({{.Initial}}r)
			{{.Initial}}h := {{.LowerEntity}}Handler.New{{.Entity}}Handler({{.Initial}}s)
`

const ServeInitializeTemplate = `
			s := server.NewRestServer(cfg.Server.Rest,`

const ServeInitializeParamsTemplate = `{{.Initial}}h, `

const ServeStartServerTemplate = `)
			err = s.StartServer()
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd
}
`

func GetServeFileConfiguration(entities []string, tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        ServeStartTemplate,
			TemplateContext: tc,
		},
		{
			Template:        ServeImportTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
		{
			Template:        ServeCommandTemplate,
			TemplateContext: tc,
		},
		{
			Template:        ServeRepositoryTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
		{
			Template:        ServeServiceAndHandlerTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
		{
			Template:        ServeInitializeTemplate,
			TemplateContext: tc,
		},
		{
			Template:        ServeInitializeParamsTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
		{
			Template:        ServeStartServerTemplate,
			TemplateContext: tc,
		},
	}
}
