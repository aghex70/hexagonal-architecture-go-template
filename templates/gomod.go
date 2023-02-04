package templates

var GoModTemplate = `module {{.ModuleName}}

go 1.19

require (
	github.com/satori/go.uuid v1.2.0
	gorm.io/gorm v1.23.8
)
`
