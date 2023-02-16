package templates

const ServeStartTemplate = `package cmd

import (
	"{{.Module}}/config"
	"{{.Module}}/persistence/database"
	"github.com/spf13/cobra"
`

const ServeImportTemplate = `	{{.LowerEntity}}Service "{{.Module}}/internal/core/services/{{.LowerEntity}}"
	{{.LowerEntity}}Handler "{{.Module}}/internal/handlers/{{.LowerEntity}}"
	{{.LowerEntity}}Repository "{{.Module}}/internal/repositories/{{.LowerEntity}}"
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

const ServeRepositoryTemplate = `			{{.Initial}}r, _ := {{.LowerEntity}}Repository.New{{.Entity}}GormRepository(gdb)
`

const ServeServiceAndHandlerTemplate = `
			{{.Initial}}s, _ := {{.LowerEntity}}Service.New{{.Entity}}Service({{.Initial}}r)
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

const importTemplate = `package cmd

import (
	"database/sql"
	"{{.Module}}/persistence/database"
	"github.com/spf13/cobra"
)

`

const makeMigrationsCommandTemplate = `func MakeMigrationsCommand(db *sql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "makemigrations [filename]",
		Short: "Generate database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				panic("fooooooooooo")
			}
			filename := args[0]
			err := database.MakeMigrations(db, filename)
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd

}

`

const migrateCommandTemplate = `func MigrateCommand(db *sql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Apply database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			err := database.Migrate(db)
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd
}

`

const rootCommandTemplate = `package cmd

import (
	"{{.Module}}/config"
	"{{.Module}}/persistence/database"
	"github.com/spf13/cobra"
)

func RootCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "{{.ProjectName}}",
		Short: "Root command",
	}

	// Intialize database
	db, err := database.NewSqlDB(*cfg.Database)
	if err != nil {
		panic(err)
	}

	cmd.AddCommand(ServeCommand(cfg))
	cmd.AddCommand(MakeMigrationsCommand(db))
	cmd.AddCommand(MigrateCommand(db))
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

func GetMakeMigrationsFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        importTemplate,
			TemplateContext: tc,
		},
		{
			Template:        makeMigrationsCommandTemplate,
			TemplateContext: tc,
		},
	}
}

func GetMigrateFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        importTemplate,
			TemplateContext: tc,
		},
		{
			Template:        migrateCommandTemplate,
			TemplateContext: tc,
		},
	}
}

func GetRootFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        rootCommandTemplate,
			TemplateContext: tc,
		},
	}
}
