package templates

var RepositoryTemplate = `package {{.LowerEntity}}

import (
	"context"
	"database/sql"
	"{{.Module}}/internal/core/domain"
	"gorm.io/gorm"
	"log"
	"time"
)

type {{.Entity}}GormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type {{.Entity}} struct {
	Id            int        ` + "`gorm:\"primaryKey;column:id\"`" + `
	Name          string     ` + "`gorm:\"column:name\"`" + `
	CreationDate  time.Time  ` + "`gorm:\"column:creation_date;autoCreateTime\"`" + `
	UpdateDate    time.Time  ` + "`gorm:\"column:update_date\"`" + `
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func ({{.Entity}}) TableName() string {
	return "{{.ProjectName}}_{{.LowerEntity}}s"
}

func (gr *{{.Entity}}GormRepository) Create(ctx context.Context, {{.Initial}} domain.{{.Entity}}) (domain.{{.Entity}}, error) {
	n{{.Initial}} := fromDto({{.Initial}})
	result := gr.DB.Create(&n{{.Initial}})
	if result.Error != nil {
		return domain.{{.Entity}}{}, result.Error
	}
	return n{{.Initial}}.ToDto(), nil
}

func (gr *{{.Entity}}GormRepository) Update(ctx context.Context, {{.Initial}} domain.{{.Entity}}) error {
	n{{.Initial}} := fromDto({{.Initial}})
	result := gr.DB.Model(&n{{.Initial}}).Where({{.Entity}}{Id: n{{.Initial}}.Id}).Updates(map[string]interface{}{
		"update_date": n{{.Initial}}.UpdateDate,
		"name": n{{.Initial}}.Name,
	})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *{{.Entity}}GormRepository) GetById(ctx context.Context, id int) (domain.{{.Entity}}, error) {
	var n{{.Initial}} {{.Entity}}
	result := gr.DB.Where(&{{.Entity}}{Id: id}).First(&n{{.Initial}})
	if result.Error != nil {
		return domain.{{.Entity}}{}, result.Error
	}
	return n{{.Initial}}.ToDto(), nil
}

func (gr *{{.Entity}}GormRepository) List(ctx context.Context) ([]domain.{{.Entity}}, error) {
	var {{.Initial}}s []{{.Entity}}
	var n{{.Initial}}s []domain.{{.Entity}}
	result := gr.DB.Find(&{{.Initial}}s)
	if result.Error != nil {
		return []domain.{{.Entity}}{}, result.Error
	}

	for _, {{.Initial}} := range {{.Initial}}s {
		{{.Initial}}{{.Initial}}:= {{.Initial}}.ToDto()
		n{{.Initial}}s = append(n{{.Initial}}s, {{.Initial}}{{.Initial}})
	}
	return n{{.Initial}}s, nil
}

func (gr *{{.Entity}}GormRepository) Delete(ctx context.Context, id int) error {
	n{{.Initial}} := {{.Entity}}{Id: id}
	result := gr.DB.Delete(&n{{.Initial}})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func New{{.Entity}}GormRepository(db *gorm.DB) (*{{.Entity}}GormRepository, error) {
	return &{{.Entity}}GormRepository{
		DB: db,
	}, nil
}

func (td {{.Entity}}) ToDto() domain.{{.Entity}} {
	return domain.{{.Entity}}{
		Name: 		  td.Name,
		CreationDate: td.CreationDate,
		Id:           td.Id,
		UpdateDate:   td.UpdateDate,
	}
}

func fromDto(td domain.{{.Entity}}) {{.Entity}} {
	return {{.Entity}}{
		Name: 		  td.Name,
		CreationDate: td.CreationDate,
		Id:           td.Id,
		UpdateDate:   td.UpdateDate,
	}
}
`

func GetRepositoryFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        RepositoryTemplate,
			TemplateContext: tc,
		},
	}
}
