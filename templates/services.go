package templates

var ServicesStartTemplate = `package ports

import (
	"context"
	"{{.Module}}/internal/core/domain"
	"net/http"
)

`

var ServicesRepeatTemplate = `type {{.Entity}}Servicer interface {
	Create(ctx context.Context, r *http.Request, req Create{{.Entity}}Request) (domain.{{.Entity}}, error)
	Update(ctx context.Context, r *http.Request, req Update{{.Entity}}Request) error
	Get(ctx context.Context, r *http.Request, req Get{{.Entity}}Request) (domain.{{.Entity}}, error)
	Delete(ctx context.Context, r *http.Request, req Delete{{.Entity}}Request) error
	List(ctx context.Context, r *http.Request) ([]domain.{{.Entity}}, error)
}

`

func GetServicesFileConfiguration(entities []string, tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        ServicesStartTemplate,
			TemplateContext: tc,
		},
		{
			Template:        ServicesRepeatTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
	}
}
