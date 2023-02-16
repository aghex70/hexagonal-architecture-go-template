package templates

const backendDockerfileTemplate = `FROM golang:1.19 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /{{.ProjectName}}/backend/
COPY . .
CMD ["air"]
`

func GetBackendDockerfileFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        backendDockerfileTemplate,
			TemplateContext: tc,
		}}
}
