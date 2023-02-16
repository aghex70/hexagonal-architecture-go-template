package templates

const mainTemplate = `package main

import (
	"{{.Module}}/cmd"
	"{{.Module}}/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	rootCmd := cmd.RootCommand(cfg)
	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}

}
`

func GetMainFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        mainTemplate,
			TemplateContext: tc,
		}}
}
