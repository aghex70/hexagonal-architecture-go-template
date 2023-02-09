package templates

var ReadmeTemplate = `#{{.ProjectName}}
{{.ProjectDescription}}

PS: This project has been generated with https://github.com/aghex70/hexagonal-architecture-go-template
`

func GetReadmeFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        ReadmeTemplate,
			TemplateContext: tc,
		}}
}
