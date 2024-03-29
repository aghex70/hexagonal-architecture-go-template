package templates

var GormRepositoryTemplate = `package {{.LowerEntity}}

import (
	"context"
	"database/sql"
	"{{.Module}}/internal/core/domain"
	"gorm.io/gorm"
	"log"
	"time"
)

type GormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type {{.Entity}} struct {
	Id            string     ` + "`gorm:\"primaryKey;column:id\"`" + `
	Name          string     ` + "`gorm:\"column:name\"`" + `
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func ({{.Entity}}) TableName() string {
	return "{{.ProjectName}}_{{.TableSuffix}}"
}

func (gr *GormRepository) Create(ctx context.Context, {{.Initial}} domain.{{.Entity}}) (domain.{{.Entity}}, error) {
	n{{.Initial}} := fromDto({{.Initial}})
	result := gr.DB.Create(&n{{.Initial}})
	if result.Error != nil {
		return domain.{{.Entity}}{}, result.Error
	}
	return n{{.Initial}}.ToDto(), nil
}

func (gr *GormRepository) Update(ctx context.Context, {{.Initial}} domain.{{.Entity}}) (domain.{{.Entity}}, error) {
	u{{.Initial}} := fromDto({{.Initial}})
	result := gr.DB.Model(&u{{.Initial}}).Where({{.Entity}}{Id: u{{.Initial}}.Id}).Updates(map[string]interface{}{
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

func (gr *GormRepository) GetById(ctx context.Context, uuid string) (domain.{{.Entity}}, error) {
	var {{.Initial}}{{.Initial}} {{.Entity}}
	result := gr.DB.Where(&{{.Entity}}{Id: uuid}).First(&{{.Initial}}{{.Initial}})
	if result.Error != nil {
		return domain.{{.Entity}}{}, result.Error
	}
	return {{.Initial}}{{.Initial}}.ToDto(), nil
}

func (gr *GormRepository) List(ctx context.Context) ([]domain.{{.Entity}}, error) {
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

func (gr *GormRepository) Delete(ctx context.Context, uuid string) error {
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

func NewGormRepository(db *gorm.DB) (*GormRepository, error) {
	return &GormRepository{
		DB: db,
	}, nil
}

func ({{.Initial}}{{.Initial}} {{.Entity}}) ToDto() domain.{{.Entity}} {
	return domain.{{.Entity}}{
		Name: 		  {{.Initial}}{{.Initial}}.Name,
		CreatedAt:    {{.Initial}}{{.Initial}}.CreatedAt,
		Id:           {{.Initial}}{{.Initial}}.Id,
		UpdatedAt:     {{.Initial}}{{.Initial}}.UpdatedAt,
	}
}

func fromDto({{.Initial}}{{.Initial}} domain.{{.Entity}}) {{.Entity}} {
	return {{.Entity}}{
		Name: 		  {{.Initial}}{{.Initial}}.Name,
		CreatedAt:    {{.Initial}}{{.Initial}}.CreatedAt,
		Id:           {{.Initial}}{{.Initial}}.Id,
		UpdatedAt:    {{.Initial}}{{.Initial}}.UpdatedAt,
	}
}
`

func GetGormRepositoryFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        GormRepositoryTemplate,
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
	case "h":
		return entity[:lastLetterIndex] + "es"
	default:
		return entity + "s"
	}
}
