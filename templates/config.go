package templates

const CacheConfigTemplate = `package config

import (
	"github.com/spf13/viper"
)

type CacheConfig struct {
	DB       int    ` + "`" + "mapstructure:\"CACHE_DB\"`" + `
	Host     string ` + "`" + "mapstructure:\"CACHE_HOST\"`" + `
	Port     int    ` + "`" + "mapstructure:\"CACHE_PORT\"`" + `
	Name     string ` + "`" + "mapstructure:\"CACHE_NAME\"`" + `
	User     string ` + "`" + "mapstructure:\"CACHE_USER\"`" + `
	Password string ` + "`" + "mapstructure:\"CACHE_PASSWORD\"`" + `
}

func LoadCacheConfig() *CacheConfig {
	cfg := &CacheConfig{}
	viper.Unmarshal(cfg)
	return cfg
}
`

const MySQLConfigTemplate = `package config

import (
	"github.com/spf13/viper"
	"time"
)

type DatabaseConfig struct {
	Dialect            string 		 ` + "`" + "mapstructure:\"DB_DIALECT\"`" + `
	Host               string 		 ` + "`" + "mapstructure:\"DB_HOST\"`" + `
	MaxOpenConnections int           ` + "`" + "mapstructure:\"DB_MAX_OPEN_CONNECTIONS\"`" + `
	MaxIdleConnections int           ` + "`" + "mapstructure:\"DB_MAX_IDLE_CONNECTIONS\"`" + `
	MaxConnLifeTime    time.Duration ` + "`" + "mapstructure:\"DB_MAX_CONN_LIFE_TIME\"`" + `
	MigrationDir       string        ` + "`" + "mapstructure:\"DB_MIGRATION_DIR\"`" + `
	Name               string        ` + "`" + "mapstructure:\"DB_NAME\"`" + `
	Net                string        ` + "`" + "mapstructure:\"DB_NETWORK\"`" + `
	Port               int           ` + "`" + "mapstructure:\"DB_PORT\"`" + `
	Password           string        ` + "`" + "mapstructure:\"DB_PASSWORD\"`" + `
	User               string        ` + "`" + "mapstructure:\"DB_USER\"`" + `
}

func LoadDatabaseConfig() *DatabaseConfig {
	cfg := &DatabaseConfig{}
	viper.Unmarshal(cfg)
	return cfg
}
`

const RestConfigTemplate = `package config

import "github.com/spf13/viper"

type RestConfig struct {
	Host string ` + "`" + "mapstructure:\"HOST\"`" + `
	Port int    ` + "`" + "mapstructure:\"PORT\"`" + `
}

func LoadRestConfig() *RestConfig {
	cfg := &RestConfig{}
	viper.Unmarshal(cfg)
	return cfg
}
`

const GRPCConfigTemplate = `package config

import "github.com/spf13/viper"

type GRPCConfig struct {
	GRPCHost string ` + "`" + "mapstructure:\"GRPC_HOST\"`" + `
	GRPCPort int    ` + "`" + "mapstructure:\"GRPC_PORT\"`" + `
}

func LoadGRPCConfig() *GRPCConfig {
	cfg := &GRPCConfig{}
	viper.Unmarshal(cfg)
	return cfg
}
`

const ServerConfigTemplate = `package config

type ServerConfig struct {
	GRPC *GRPCConfig
	Rest *RestConfig
}

func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		GRPC: LoadGRPCConfig(),
		Rest: LoadRestConfig(),
	}
}
`

func GetCacheFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        CacheConfigTemplate,
			TemplateContext: tc,
		},
	}
}

func GetMySQLFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        MySQLConfigTemplate,
			TemplateContext: tc,
		},
	}
}

func GetRestFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        RestConfigTemplate,
			TemplateContext: tc,
		},
	}
}

func GetGRPCFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        GRPCConfigTemplate,
			TemplateContext: tc,
		},
	}
}

func GetServerFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        ServerConfigTemplate,
			TemplateContext: tc,
		},
	}
}
