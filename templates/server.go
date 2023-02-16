package templates

const serverTemplate = `package server

type Server interface {
	StartServer() error
	StopServer() error
}
`

func GetServerInterfaceFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        serverTemplate,
			TemplateContext: tc,
		},
	}
}
