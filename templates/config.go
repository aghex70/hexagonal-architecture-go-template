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
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil
	}
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
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil
	}
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
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil
	}
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
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil
	}
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

const ConfigConfigTemplate = `package config

import (
	"github.com/spf13/viper"
)

const (
	CONFIG_NAME string = ".env"
	CONFIG_TYPE string = "env"
	CONFIG_PATH string = "./"
)

type Config struct {
	Cache     *CacheConfig
	Database  *DatabaseConfig
	Server    *ServerConfig
}

func NewConfig() (*Config, error) {
	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &Config{
		Cache:     LoadCacheConfig(),
		Database:  LoadDatabaseConfig(),
		Server:    LoadServerConfig(),
	}, nil
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

func GetConfigFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        ConfigConfigTemplate,
			TemplateContext: tc,
		},
	}
}
