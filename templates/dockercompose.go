package templates

var dockerComposeBaseTemplate = `version: '3.9'

services:
  backend:
    build:
      context: backend
      dockerfile: Dockerfile
    command: ["air", "migrate", "&&", "air", "serve"]
    env_file:
      - backend/.env
    ports:
      - "11001:11001"
    depends_on:
      db:
        condition: service_healthy

`

var dockerComposeFrontendTemplate = `  frontend:
    build:
      context: frontend
      dockerfile: Dockerfile
    command: ["npm", "run", "start:prod"]
    env_file:
      - frontend/.prod.env
    ports:
      - "3000:3000"
    stdin_open: true

`

var dockerComposeMySqlTemplate = `  db:
    image: mysql:latest
    environment:
      - MYSQL_ROOT_PASSWORD=root
      - MYSQL_DATABASE={{.ProjectName}}db
      - MYSQL_USER={{.ProjectName}}user
      - MYSQL_PASSWORD=root
      - MYSQL_TCP_PORT=11306
    ports:
      - "11306:11306"
    healthcheck:
      test: [ "CMD", "mysqladmin", "ping", "-h", "db", "-u$$MYSQL_USER", "-p$$MYSQL_ROOT_PASSWORD" ]
      interval: 5s
      timeout: 5s
      retries: 10

`

var dockerComposePostgreSQLTemplate = `  db:
    image: postgres

`

var dockerComposeMongoDBTemplate = `  db:
    image: bitnami/mongodb
    environment:
        - MONGO_INITDB_DATABASE={{.ProjectName}}db
    ports:
      - "27017:27017"
    healthcheck:
      test: echo 'db.runCommand("ping").ok' | mongo db:27017/{{.ProjectName}}db --quiet
      interval: 5s
      timeout: 5s
      retries: 10

`

var dockerComposeRedisTemplate = `  redis:
    image: redis:alpine
    command: [ redis-server, --port, "10379" ]

`

func GetDockerComposeFileConfiguration(ec ExtraConfiguration, tc TemplateContext) []FileConfiguration {
	fc := []FileConfiguration{
		{
			Template:        dockerComposeBaseTemplate,
			TemplateContext: tc,
		},
	}
	if ec.Frontend {
		fc = append(fc, FileConfiguration{
			Template:        dockerComposeFrontendTemplate,
			TemplateContext: tc,
		})
	}
	if ec.MySQL {
		fc = append(fc, FileConfiguration{
			Template:        dockerComposeMySqlTemplate,
			TemplateContext: tc,
		})
	}
	if ec.Postgres {
		fc = append(fc, FileConfiguration{
			Template:        dockerComposePostgreSQLTemplate,
			TemplateContext: tc,
		})
	}
	if ec.MongoDB {
		fc = append(fc, FileConfiguration{
			Template:        dockerComposeMongoDBTemplate,
			TemplateContext: tc,
		})
	}
	if ec.Redis {
		fc = append(fc, FileConfiguration{
			Template:        dockerComposeRedisTemplate,
			TemplateContext: tc,
		})
	}
	return fc
}
