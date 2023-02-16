package templates

const baseBackendEnvTemplate = `# Server
SERVER_HOST=backend
SERVER_PORT=11001

# Database
DB_NAME={{.ProjectName}}db
DB_HOST=db
DB_PORT=11306
DB_USER={{.ProjectName}}user
DB_PASSWORD=root
DB_MAX_OPEN_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=10
DB_MAX_CONNECTION_LIFE_TIME=0
DB_DIALECT=mysql
DB_MIGRATION_DIR=migrations
DB_NETWORK=tcp
`

const cacheEnvTemplate = `
# Cache
CACHE_NAME=redis
CACHE_HOST=redis
CACHE_PORT=11379
CACHE_USER=redis
CACHE_PASSWORD=redis
CACHE_DB=0
`

func GetBackendEnvFileConfiguration(ec ExtraConfiguration, tc TemplateContext) []FileConfiguration {
	fc := []FileConfiguration{
		{
			Template:        baseBackendEnvTemplate,
			TemplateContext: tc,
		},
	}

	if ec.Redis {
		fc = append(fc, FileConfiguration{
			Template:        cacheEnvTemplate,
			TemplateContext: tc,
		})
	}
	return fc
}
