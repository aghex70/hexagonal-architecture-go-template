package templates

type DomainData struct {
	Entity      string
	LowerEntity string
	Module      string
	ProjectName string
}

var DomainTemplate = `package domain

import "time"

type {{.Entity}} struct {
	Id				int			` + "`json:\"id\"`" + `
	Name			string		` + "`json:\"name\"`" + `
	CreationDate	time.Time	` + "`json:\"creationDate\"`" + `
	UpdateDate		time.Time	` + "`json:\"updateDate\"`" + `
}
`
