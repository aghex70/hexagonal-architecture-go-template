package templates

var RepositoryImportTemplate = `package ports

import (
	"context"
	"{{.Module}}/internal/core/domain"
)

`

var RepositoryRepeatTemplate = `type {{.Entity}}Repository interface {
	Create(ctx context.Context, a domain.{{.Entity}}) (domain.{{.Entity}}, error)
	Update(ctx context.Context, a domain.{{.Entity}}) (domain.{{.Entity}}, error)
	GetById(ctx context.Context, uuid string) (domain.{{.Entity}}, error)
	List(ctx context.Context) ([]domain.{{.Entity}}, error)
	Delete(ctx context.Context, uuid string) error
}

`

func GetRepositoryFileConfiguration(entities []string, tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        RepositoryImportTemplate,
			TemplateContext: tc,
		},
		{
			Template:        RepositoryRepeatTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
	}
}
