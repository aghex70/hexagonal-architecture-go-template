package hexagonal_architecture_go_template

import (
	"github.com/aghex70/hexagonal-architecture-go-template/templates"
	"os"
)

var projectBasePath string

func main() {
	projectBasePath := os.Args[1]
	pe, err := projectExists(projectBasePath)
	if err != nil {
		panic(err)
	}

	if pe == true {
		panic(ProjectFileError)
	}

	projectName, err := scanString(ScanProjectName)
	if err != nil {
		panic(err)
	}

	projectPath := generateProjectPath(projectBasePath, projectName)
	err = createDirectory(projectPath)

	version, err := scanStringWithDefault(ScanVersion, DefaultVersion)
	if err != nil {
		panic(err)
	}

	description, err := scanStringWithDefault(ScanDescription, DefaultDescription)
	if err != nil {
		panic(err)
	}

	frontend, err := scanStringCastBoolean(ScanReact, DefaultActive)
	if err != nil {
		panic(err)
	}

	sequel, err := scanStringCastBoolean(ScanSqlDatabase, DefaultActive)
	if err != nil {
		panic(err)
	}

	mysql, err := scanStringCastBoolean(ScanMysql, DefaultInactive)
	if err != nil {
		panic(err)
	}

	postgres, err := scanStringCastBoolean(ScanPostgres, DefaultActive)
	if err != nil {
		panic(err)
	}

	nosequel, err := scanStringCastBoolean(ScanNoSqlDatabase, DefaultActive)
	if err != nil {
		panic(err)
	}

	mongodb, err := scanStringCastBoolean(ScanMongodb, DefaultActive)
	if err != nil {
		panic(err)
	}

	redis, err := scanStringCastBoolean(ScanRedis, DefaultActive)
	if err != nil {
		panic(err)
	}

	nginx, err := scanStringCastBoolean(ScanReverseProxy, DefaultInactive)
	if err != nil {
		panic(err)
	}

	api, err := scanStringCastBoolean(ScanApi, DefaultActive)
	if err != nil {
		panic(err)
	}

	rest, err := scanStringCastBoolean(ScanRest, DefaultActive)
	if err != nil {
		panic(err)
	}

	grpc, err := scanStringCastBoolean(ScanGrpc, DefaultActive)
	if err != nil {
		panic(err)
	}

	graphql, err := scanStringCastBoolean(ScanGraphql, DefaultInactive)
	if err != nil {
		panic(err)
	}

	entities, err := scanMultipleStrings(ScanEntities)
	if err != nil {
		panic(err)
	}

	err = generateStubs(projectPath, BackendDirectory)
	if len(entities) > 0 {

	}
	if frontend == true {
		err := generateStubs(projectPath, FrontendDirectory)
		if err != nil {
			panic(err)
		}
	}

	err = generateFile(projectPath, BackendDirectory, entities[0], templates.DomainTemplate, templates.DomainData)
	if err != nil {
		panic(err)
	}
}
