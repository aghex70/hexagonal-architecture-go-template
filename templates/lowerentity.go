package templates

var LowerEntityTemplate = `package {{.LowerEntity}}`

func GetLowerEntityFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{{
		Template:        LowerEntityTemplate,
		TemplateContext: tc,
	},
	}
}
