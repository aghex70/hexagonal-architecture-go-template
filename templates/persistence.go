package templates

const mySQLTemplate = `package database

import (
	"database/sql"
	"fmt"
	"{{.Module}}/config"
	"github.com/go-sql-driver/mysql"
)

func NewSqlDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	mySqlConfig := mysql.NewConfig()
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	mySqlConfig.Addr = address
	mySqlConfig.DBName = cfg.Name
	mySqlConfig.User = cfg.User
	mySqlConfig.Passwd = cfg.Password
	mySqlConfig.Net = cfg.Net
	mySqlConfig.ParseTime = true

	sqlDB, err := sql.Open(cfg.Dialect, mySqlConfig.FormatDSN())
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(cfg.MaxConnLifeTime)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	return sqlDB, nil
}
`

const gormTemplate = `package database

import (
	"fmt"
	"{{.Module}}/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Net, cfg.Host, cfg.Port, cfg.Name)
	//dsn = "{{.ProjectName}}user:{{.ProjectName}}pw@tcp(db:11306)/{{.ProjectName}}"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
`

const migrationTemplate = `package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/pressly/goose"
)

var (
	ErrMigrationFile = errors.New("please provide a valid name for the migration file")
)

const (
	// Migration directory
	migrationDirectory string = "persistence/database/migrations"

	// Migrations table
	migrationTable string = "{{.ProjectName}}_db_version"
)

func init() {
	goose.SetTableName(migrationTable)
}

func Migrate(db *sql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Run("up", db, migrationDirectory, "sql"); err != nil {
		fmt.Printf("%+v", err)
		return err
	}

	return nil
}

func MakeMigrations(db *sql.DB, filename string) error {
	if filename == "" {
		return ErrMigrationFile
	}

	if err := goose.Run("create", db, migrationDirectory, filename, "sql"); err != nil {
		return err
	}

	return nil
}
`

func GetMySQLPersistenceFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        mySQLTemplate,
			TemplateContext: tc,
		}}
}

func GetGormPersistenceFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        gormTemplate,
			TemplateContext: tc,
		}}
}

func GetMigrationPersistenceFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        migrationTemplate,
			TemplateContext: tc,
		}}
}
