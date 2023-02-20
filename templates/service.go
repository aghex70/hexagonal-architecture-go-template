package templates

var ServiceTemplate = `package {{.LowerEntity}}

import (
	"context"
	"{{.Module}}/internal/core/domain"
	"{{.Module}}/internal/core/ports"
	"{{.Module}}/internal/stores"
	"time"
)

type {{.Entity}}Service struct {
	{{.LowerEntity}}Repository   stores.{{.Entity}}Repository
}

func (s {{.Entity}}Service) Create(ctx context.Context, r ports.Create{{.Entity}}Request) (domain.{{.Entity}}, error) {
	{{.Initial}}{{.Initial}} := domain.{{.Entity}}{
		CreationDate: time.Now(),
		Name: 		  r.Name,
	}
	ne, err := s.{{.LowerEntity}}Repository.Create(ctx, {{.Initial}}{{.Initial}})
	if err != nil {
		return domain.{{.Entity}}{}, err
	}

	return ne, nil
}

func (s {{.Entity}}Service) Update(ctx context.Context, r ports.Update{{.Entity}}Request) (domain.{{.Entity}}, error) {
	{{.Initial}}{{.Initial}} := domain.{{.Entity}}{
		Id:                 r.Id,
		UpdateDate: 		time.Now(),
		Name: 		  		r.Name,
	}
	u{{.Initial}}, err := s.{{.LowerEntity}}Repository.Update(ctx, {{.Initial}}{{.Initial}})
	if err != nil {
		return domain.{{.Entity}}{}, err
	}
	return u{{.Initial}}, nil
}

func (s {{.Entity}}Service) Get(ctx context.Context, uuid string) (domain.{{.Entity}}, error) {
	{{.Initial}}{{.Initial}}, err := s.{{.LowerEntity}}Repository.GetById(ctx, uuid)
	if err != nil {
		return domain.{{.Entity}}{}, err
	}
	return {{.Initial}}{{.Initial}}, nil
}

func (s {{.Entity}}Service) Delete(ctx context.Context, uuid string) error {
	err := s.{{.LowerEntity}}Repository.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s {{.Entity}}Service) List(ctx context.Context) ([]domain.{{.Entity}}, error) {
	{{.Initial}}s, err := s.{{.LowerEntity}}Repository.List(ctx)
	if err != nil {
		return []domain.{{.Entity}}{}, err
	}
	return {{.Initial}}s, nil
}

func New{{.Entity}}Service({{.Initial}}r stores.{{.Entity}}Repository) ({{.Entity}}Service, error) {
	return {{.Entity}}Service{
		{{.LowerEntity}}Repository:      {{.Initial}}r,
	}, nil
}
`

func GetServiceFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        ServiceTemplate,
			TemplateContext: tc,
		},
	}
}
