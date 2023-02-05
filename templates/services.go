package templates

var ServicesStartTemplate = `package ports

import (
	"context"
	"{{.Module}}/internal/core/domain"
	"net/http"
)

`

var ServicesEntitiesTemplate = `type {{.Entity}}Servicer interface {
	Create(ctx context.Context, r *http.Request, req Create{{.Entity}}Request) (domain.{{.Entity}}, error)
	Update(ctx context.Context, r *http.Request, req Update{{.Entity}}Request) error
	Get(ctx context.Context, r *http.Request, req Get{{.Entity}}Request) (domain.{{.Entity}}, error)
	Delete(ctx context.Context, r *http.Request, req Delete{{.Entity}}Request) error
	List(ctx context.Context, r *http.Request) ([]domain.{{.Entity}}, error)
}

`
