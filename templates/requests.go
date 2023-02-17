package templates

var RequestsStartTemplate = `package ports

import "time"

`

var RequestsRepeatTemplate = `type Create{{.Entity}}Request struct {
	CreationDate    time.Time 	` + "`json:\"creationDate\"`" + `
	Name       		string 		` + "`json:\"name\"`" + `
}

type Update{{.Entity}}Request struct {
	Id					string  	` + "`json:\"id\"`" + `
	Name       			string 		` + "`json:\"name\"`" + `
	UpdateDate        	time.Time 	` + "`json:\"updateDate\"`" + `
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
