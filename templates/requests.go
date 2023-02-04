package templates

var requestsTemplate = `package ports

type Create{{.Entity}}Request struct {
	CreationDate       time.Time ` + "`json:\"creationDate\"`" + `
}

type Update{{.Entity}}Request struct {
	Id					int  ` + "`json:\"id\"`" + `
	UpdateDate        	time.Time ` + "`json:\"updateDate\"`" + `
}

type Get{{.Entity}}Request struct {
	Id	int  ` + "`json:\"id\"`" + `
}

type Delete{{.Entity}}Request struct {
	Id	int  ` + "`json:\"id\"`" + `
}
`
