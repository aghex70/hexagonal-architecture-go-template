package templates

var ServiceTemplate = `package {{.LowerEntity}}

import (
	"context"
	"{{.Module}}/internal/core/domain"
	"{{.Module}}/internal/core/ports"
	{{.LowerEntity}}Store "{{.Module}}/internal/stores/{{.LowerEntity}}"
)

type Service struct {
	{{.LowerEntity}}Repository   {{.LowerEntity}}Store.{{.Entity}}GormRepository
}

func (s Service) Create(ctx context.Context, r ports.Create{{.Entity}}Request) (domain.{{.Entity}}, error) {
	{{.Initial}}{{.Initial}} := domain.{{.Entity}}{
		Name: 		  r.Name,
	}
	ne, err := s.{{.LowerEntity}}Repository.Create(ctx, {{.Initial}}{{.Initial}})
	if err != nil {
		return domain.{{.Entity}}{}, err
	}

	return ne, nil
}

func (s Service) Update(ctx context.Context, r ports.Update{{.Entity}}Request) (domain.{{.Entity}}, error) {
	{{.Initial}}{{.Initial}} := domain.{{.Entity}}{
		Id:                 r.Id,
		Name: 		  		r.Name,
	}
	u{{.Initial}}, err := s.{{.LowerEntity}}Repository.Update(ctx, {{.Initial}}{{.Initial}})
	if err != nil {
		return domain.{{.Entity}}{}, err
	}
	return u{{.Initial}}, nil
}

func (s Service) Get(ctx context.Context, uuid string) (domain.{{.Entity}}, error) {
	{{.Initial}}{{.Initial}}, err := s.{{.LowerEntity}}Repository.GetById(ctx, uuid)
	if err != nil {
		return domain.{{.Entity}}{}, err
	}
	return {{.Initial}}{{.Initial}}, nil
}

func (s Service) Delete(ctx context.Context, uuid string) error {
	err := s.{{.LowerEntity}}Repository.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context) ([]domain.{{.Entity}}, error) {
	{{.Initial}}s, err := s.{{.LowerEntity}}Repository.List(ctx)
	if err != nil {
		return []domain.{{.Entity}}{}, err
	}
	return {{.Initial}}s, nil
}

func NewService({{.Initial}}r {{.LowerEntity}}Store.{{.Entity}}GormRepository) (Service, error) {
	return Service{
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
