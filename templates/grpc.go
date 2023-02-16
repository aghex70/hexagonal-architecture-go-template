package templates

const grpcTemplate = `package server`

func GetGRPCInterfaceFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        grpcTemplate,
			TemplateContext: tc,
		},
	}
}
