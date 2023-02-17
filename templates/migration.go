package templates

const MigrationStartTemplate = `-- +goose Up
`

const MigrationRepeatTemplate = `CREATE TABLE {{.ProjectName}}_{{.TableSuffix}} (
    id VARCHAR(36) NOT NULL PRIMARY KEY DEFAULT uuid(),
    name VARCHAR(128),
    creation_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_date TIMESTAMP NULL,
);

`

func GetMigrationFileConfiguration(entities []string, tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        MigrationStartTemplate,
			TemplateContext: tc,
		},
		{
			Template:        MigrationRepeatTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		},
	}
}
