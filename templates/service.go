package templates

var ServiceTemplate = `package {{.LowerEntity}}

import (
	"context"
	"{{.Module}}/internal/core/domain"
	"{{.Module}}/internal/core/ports"
	"{{.Module}}/internal/repositories/{{.LowerEntity}}"
	"{{.Module}}/server"
	"log"
	"net/http"
	"strconv"
	"time"
)

type {{.Entity}}Service struct {
	logger                 	*log.Logger
	{{.LowerEntity}}Repository   *{{.Entity}}.{{.Entity}}GormRepository
}

func (s {{.Entity}}Service) Create(ctx context.Context, r *http.Request, req ports.Create{{.Entity}}Request) (domain.{{.Entity}}, error) {
	e := domain.{{.Entity}}{
		CreationDate: time.Now(),
		Name: 		  req.Name,
	}
	ne, err := s.{{.LowerEntity}}Repository.Create(ctx, e)
	if err != nil {
		return domain.{{.Entity}}{}, err
	}

	return ne, nil
}

func (s {{.Entity}}Service) Update(ctx context.Context, r *http.Request, req ports.Update{{.Entity}}Request) error {
	e := domain.{{.Entity}}{
		Id:                 int(req.{{.Entity}}Id),
		UpdateDate: 		time.Now(),
		Name: 		  		req.Name,
	}
	err := s.{{.LowerEntity}}Repository.Update(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

func (s {{.Entity}}Service) Get(ctx context.Context, r *http.Request, req ports.Get{{.Entity}}Request) (domain.{{.Entity}}, error) {
	e, err := s.{{.LowerEntity}}Repository.GetById(ctx, req.Id)
	if err != nil {
		return domain.{{.Entity}}{}, err
	}
	return e, nil
}

func (s {{.Entity}}Service) Delete(ctx context.Context, r *http.Request, req ports.Delete{{.Entity}}Request) error {
	err := s.{{.LowerEntity}}Repository.Delete(ctx, int(req.{{.Entity}}Id), int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s {{.Entity}}Service) List(ctx context.Context, r *http.Request) ([]domain.{{.Entity}}, error) {
	es, err := s.{{.LowerEntity}}Repository.List(ctx)
	if err != nil {
		return []domain.{{.Entity}}{}, err
	}
	return es, nil
}

func New{{.Entity}}Service(er *{{.Entity}}.{{.Entity}}GormRepository, logger *log.Logger) {{.Entity}}Service {
	return {{.Entity}}Service{
		logger:                 	logger,
		{{.LowerEntity}}Repository:      er,
	}
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
