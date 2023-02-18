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
	Id            string     ` + "`gorm:\"primaryKey;column:id\"`" + `
	Name          string     ` + "`gorm:\"column:name\"`" + `
	CreationDate  time.Time  ` + "`gorm:\"column:creation_date;autoCreateTime\"`" + `
	UpdateDate    time.Time  ` + "`gorm:\"column:update_date\"`" + `
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func ({{.Entity}}) TableName() string {
	return "{{.ProjectName}}_{{.TableSuffix}}"
}

func (gr *{{.Entity}}GormRepository) Create(ctx context.Context, {{.Initial}} domain.{{.Entity}}) (domain.{{.Entity}}, error) {
	n{{.Initial}} := fromDto({{.Initial}})
	result := gr.DB.Create(&n{{.Initial}})
	if result.Error != nil {
		return domain.{{.Entity}}{}, result.Error
	}
	return n{{.Initial}}.ToDto(), nil
}

func (gr *{{.Entity}}GormRepository) Update(ctx context.Context, {{.Initial}} domain.{{.Entity}}) (domain.{{.Entity}}, error) {
	u{{.Initial}} := fromDto({{.Initial}})
	result := gr.DB.Model(&u{{.Initial}}).Where({{.Entity}}{Id: u{{.Initial}}.Id}).Updates(map[string]interface{}{
		"update_date": u{{.Initial}}.UpdateDate,
		"name": u{{.Initial}}.Name,
	})

	if result.RowsAffected == 0 {
		return domain.{{.Entity}}{}, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return domain.{{.Entity}}{}, result.Error
	}
	return u{{.Initial}}.ToDto(), nil
}

func (gr *{{.Entity}}GormRepository) GetById(ctx context.Context, uuid string) (domain.{{.Entity}}, error) {
	var {{.Initial}}{{.Initial}} {{.Entity}}
	result := gr.DB.Where(&{{.Entity}}{Id: uuid}).First(&{{.Initial}}{{.Initial}})
	if result.Error != nil {
		return domain.{{.Entity}}{}, result.Error
	}
	return {{.Initial}}{{.Initial}}.ToDto(), nil
}

func (gr *{{.Entity}}GormRepository) List(ctx context.Context) ([]domain.{{.Entity}}, error) {
	var {{.Initial}}s []{{.Entity}}
	var {{.Initial}}{{.Initial}}s []domain.{{.Entity}}
	result := gr.DB.Find(&{{.Initial}}s)
	if result.Error != nil {
		return []domain.{{.Entity}}{}, result.Error
	}

	for _, {{.Initial}} := range {{.Initial}}s {
		{{.Initial}}{{.Initial}} := {{.Initial}}.ToDto()
		{{.Initial}}{{.Initial}}s = append({{.Initial}}{{.Initial}}s, {{.Initial}}{{.Initial}})
	}
	return {{.Initial}}{{.Initial}}s, nil
}

func (gr *{{.Entity}}GormRepository) Delete(ctx context.Context, uuid string) error {
	{{.Initial}}{{.Initial}} := {{.Entity}}{Id: uuid}
	result := gr.DB.Delete(&{{.Initial}}{{.Initial}})
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

func ({{.Initial}}{{.Initial}} {{.Entity}}) ToDto() domain.{{.Entity}} {
	return domain.{{.Entity}}{
		Name: 		  {{.Initial}}{{.Initial}}.Name,
		CreationDate: {{.Initial}}{{.Initial}}.CreationDate,
		Id:           {{.Initial}}{{.Initial}}.Id,
		UpdateDate:   {{.Initial}}{{.Initial}}.UpdateDate,
	}
}

func fromDto({{.Initial}}{{.Initial}} domain.{{.Entity}}) {{.Entity}} {
	return {{.Entity}}{
		Name: 		  {{.Initial}}{{.Initial}}.Name,
		CreationDate: {{.Initial}}{{.Initial}}.CreationDate,
		Id:           {{.Initial}}{{.Initial}}.Id,
		UpdateDate:   {{.Initial}}{{.Initial}}.UpdateDate,
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

func GenerateTableSuffix(entity string) string {
	lastLetterIndex := len(entity) - 1
	lastLetter := entity[lastLetterIndex:]
	switch lastLetter {
	case "s":
		return entity
	case "y":
		return entity[:lastLetterIndex] + "ies"
	default:
		return entity + "s"
	}
}
