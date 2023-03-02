package templates

var RequestsStartTemplate = `package ports
`

var RequestsRepeatTemplate = `
type Create{{.Entity}}Request struct {
	Name       		string 		` + "`json:\"name\"`" + `
}

type Update{{.Entity}}Request struct {
	Id					string  	` + "`json:\"id\"`" + `
	Name       			string 		` + "`json:\"name\"`" + `
}
`

func GetRequestsFileConfiguration(entities []string, tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        RequestsStartTemplate,
			TemplateContext: tc,
		},
		{
			Template:        RequestsRepeatTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
	}
}
