package templates

var RepositoryTemplate = `package {{.LowerEntity}}

import (
	"context"
	"database/sql"
	"fmt"
	"{{.Module}}/internal/core/domain"
	"gorm.io/gorm"
	"log"
	"strconv"
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

func (gr *{{.Entity}}GormRepository) Create(ctx context.Context, e domain.{{.Entity}}) error {
	ne := fromDto(e)
	result := gr.DB.Create(&ne)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *{{.Entity}}GormRepository) Update(ctx context.Context, e domain.{{.Entity}}) error {
	ne := fromDto(e)
	result := gr.DB.Model(&ne).Where({{.Entity}}{Id: ne.Id}).Updates(map[string]interface{}{
		"update_date": ne.UpdateDate,
		"name": ne.Name,
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
	var ne {{.Entity}}
	result := gr.DB.Where(&{{.Entity}}{Id: id}).Find(&ne).First()
	if result.Error != nil {
		return domain.{{.Entity}}Info{}, result.Error
	}
	return ne.ToDto(), nil
}

func (gr *{{.Entity}}GormRepository) List(ctx context.Context, categoryId int) ([]domain.{{.Entity}}, error) {
	var es []{{.Entity}}
	var nes []domain.{{.Entity}}
	result := gr.DB.Find(&es)
	if result.Error != nil {
		return []domain.{{.Entity}}{}, result.Error
	}

	for _, e := range es {
		ee:= e.ToDto()
		nes = append(nes, ee)
	}
	return nes, nil
}

func (gr *{{.Entity}}GormRepository) Delete(ctx context.Context, id int) error {
	ne := {{.Entity}}{Id: id}
	result := gr.DB.Delete(&ne)
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
	return []FileConfiguration{{
		Template:        RepositoryTemplate,
		TemplateContext: tc,
	},
	}
}
