package templates

var domainTemplate = `package domain

import "time"

type {{.Entity}} struct {
	Id				int			` + "`json:\"id\"`" + `
	CreationDate	time.Time	` + "`json:\"creationDate\"`" + `
	UpdateDate		time.Time	` + "`json:\"updateDate\"`" + `
}
`
