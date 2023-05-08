package templates

var GoModTemplate = `module {{.Module}}

go 1.19

require (
	github.com/gin-gonic/gin v1.9.0
	github.com/satori/go.uuid v1.2.0
	gorm.io/gorm v1.23.8
)
`

func GetGoModFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        GoModTemplate,
			TemplateContext: tc,
		},
	}
}
