package templates

var DomainTemplate = `package domain

import "time"

type {{.Entity}} struct {
	Id				string		` + "`json:\"id\"`" + `
	Name			string		` + "`json:\"name\"`" + `
	CreatedAt		time.Time	` + "`json:\"createdAt\"`" + `
	UpdatedAt		time.Time	` + "`json:\"updatedAt\"`" + `
}
`

func GetDomainFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{{
		Template:        DomainTemplate,
		TemplateContext: tc,
	},
	}
}
