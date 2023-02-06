package main

import (
	"fmt"
	"github.com/aghex70/hexagonal-architecture-go-template/common"
	"github.com/aghex70/hexagonal-architecture-go-template/templates"
	"os"
	"strings"
	"time"
)

func main() {
	projectBasePath := os.Args[1]
	//pe, err := common.ProjectExists(projectBasePath)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if pe == true {
	//	panic(common.ProjectFileError)
	//}

	projectModule, err := common.ScanString(common.ScanProjectModule)
	if err != nil {
		panic(err)
	}

	projectName, err := common.ScanString(common.ScanProjectName)
	if err != nil {
		panic(err)
	}

	////version, err := scanStringWithDefault(ScanVersion, DefaultVersion)
	//_, err = common.ScanStringWithDefault(common.ScanVersion, common.DefaultVersion)
	//if err != nil {
	//	panic(err)
	//}
	//
	////description, err := scanStringWithDefault(ScanDescription, DefaultDescription)
	//_, err = common.ScanStringWithDefault(common.ScanDescription, common.DefaultDescription)
	//if err != nil {
	//	panic(err)
	//}
	//
	//frontend, err := common.ScanStringCastBoolean(common.ScanReact, common.DefaultActive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////sequel, err := scanStringCastBoolean(ScanSqlDatabase, DefaultActive)
	//_, err = common.ScanStringCastBoolean(common.ScanSqlDatabase, common.DefaultActive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////mysql, err := scanStringCastBoolean(ScanMysql, DefaultInactive)
	//_, err = common.ScanStringCastBoolean(common.ScanMysql, common.DefaultInactive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////postgres, err := scanStringCastBoolean(ScanPostgres, DefaultActive)
	//_, err = common.ScanStringCastBoolean(common.ScanPostgres, common.DefaultActive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////nosequel, err := scanStringCastBoolean(ScanNoSqlDatabase, DefaultActive)
	//_, err = common.ScanStringCastBoolean(common.ScanNoSqlDatabase, common.DefaultActive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////mongodb, err := scanStringCastBoolean(ScanMongodb, DefaultActive)
	//_, err = common.ScanStringCastBoolean(common.ScanMongodb, common.DefaultActive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////redis, err := scanStringCastBoolean(ScanRedis, DefaultActive)
	//_, err = common.ScanStringCastBoolean(common.ScanRedis, common.DefaultActive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////nginx, err := scanStringCastBoolean(ScanReverseProxy, DefaultInactive)
	//_, err = common.ScanStringCastBoolean(common.ScanReverseProxy, common.DefaultInactive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////api, err := scanStringCastBoolean(ScanApi, DefaultActive)
	//_, err = common.ScanStringCastBoolean(common.ScanApi, common.DefaultActive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////rest, err := scanStringCastBoolean(ScanRest, DefaultActive)
	//_, err = common.ScanStringCastBoolean(common.ScanRest, common.DefaultActive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////grpc, err := scanStringCastBoolean(ScanGrpc, DefaultActive)
	//_, err = common.ScanStringCastBoolean(common.ScanGrpc, common.DefaultActive)
	//if err != nil {
	//	panic(err)
	//}
	//
	////graphql, err := scanStringCastBoolean(ScanGraphql, DefaultInactive)
	//_, err = common.ScanStringCastBoolean(common.ScanGraphql, common.DefaultInactive)
	//if err != nil {
	//	panic(err)
	//}

	//entities, err := common.ScanMultipleStrings(common.ScanEntities)
	entities := []string{"User", "Admin", "Car"}
	if err != nil {
		panic(err)
	}

	projectPath := common.GenerateProjectPath(projectBasePath, projectName)
	err = common.CreateDirectory(projectPath)
	if err != nil {
		panic(err)
	}

	err = common.GenerateStubs(projectPath, common.BackendDirectory)
	if len(entities) > 0 {
		for _, e := range entities {
			le := strings.ToLower(e)
			tc := templates.DomainData{
				Entity:      e,
				LowerEntity: le,
				Module:      projectModule,
				ProjectName: projectName,
			}
			// Domain
			domainPath := projectPath + common.BackendDirectory + common.DomainDirectory + le
			fmt.Println("domainPath: ", domainPath)
			domainFileConfiguration := common.FileConfiguration{
				TemplateStart:   templates.DomainTemplate,
				TemplateContext: tc,
			}
			err = common.GenerateFile(domainPath, common.GolangFileExtension, domainFileConfiguration)
			if err != nil {
				panic(err)
			}

			// Repositories
			repositoryPath := projectPath + common.BackendDirectory + common.RepositoriesDirectory + le
			err = common.CreateDirectory(repositoryPath)
			if err != nil {
				panic(err)
			}

			repositoryFileConfiguration := common.FileConfiguration{
				TemplateStart:   templates.RepositoryTemplate,
				TemplateContext: tc,
			}
			err = common.GenerateFile(repositoryPath+"/"+common.GormFileName, common.GolangFileExtension, repositoryFileConfiguration)
			if err != nil {
				panic(err)
			}

			// Services
			servicePath := projectPath + common.BackendDirectory + common.ServicesDirectory + le
			err = common.CreateDirectory(servicePath)
			if err != nil {
				panic(err)
			}

			serviceFileConfiguration := common.FileConfiguration{
				TemplateStart:   templates.ServiceStartTemplate,
				TemplateEnd:     templates.ServiceEndTemplate,
				TemplateContext: tc,
			}
			err = common.GenerateFile(servicePath+"/"+common.ServiceFileName, common.GolangFileExtension, serviceFileConfiguration)
			if err != nil {
				panic(err)
			}

			serviceTestFileConfiguration := common.FileConfiguration{
				TemplateStart:   templates.LowerEntityTemplate,
				TemplateContext: tc,
			}
			err = common.GenerateFile(servicePath+"/"+common.ServiceTestFileName, common.GolangFileExtension, serviceTestFileConfiguration)
			if err != nil {
				panic(err)
			}

			validatorsFileConfiguration := common.FileConfiguration{
				TemplateStart:   templates.LowerEntityTemplate,
				TemplateContext: tc,
			}
			err = common.GenerateFile(servicePath+"/"+common.ServiceValidatorsFileName, common.GolangFileExtension, validatorsFileConfiguration)
			if err != nil {
				panic(err)
			}
		}

		tc := templates.DomainData{
			Module:      projectModule,
			ProjectName: projectName,
		}

		// Requests
		requestsFileConfiguration := common.FileConfiguration{
			TemplateStart:   templates.RequestsStartTemplate,
			TemplateRepeat:  templates.RequestsRepeatTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		}

		requestsPath := projectPath + common.BackendDirectory + common.PortsDirectory + common.RequestsFileName
		err = common.GenerateFile(requestsPath, common.GolangFileExtension, requestsFileConfiguration)
		if err != nil {
			panic(err)
		}

		// Services
		servicesFileConfiguration := common.FileConfiguration{
			TemplateStart:   templates.ServicesStartTemplate,
			TemplateRepeat:  templates.ServicesRepeatTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		}

		servicesPath := projectPath + common.BackendDirectory + common.PortsDirectory + common.ServicesFileName
		err = common.GenerateFile(servicesPath, common.GolangFileExtension, servicesFileConfiguration)
		if err != nil {
			panic(err)
		}

		// Migrations
		migrationsFileConfiguration := common.FileConfiguration{
			TemplateStart:   templates.MigrationStartTemplate,
			TemplateRepeat:  templates.MigrationRepeatTemplate,
			TemplateContext: tc,
			Repeat:          true,
			RepeatEntities:  entities,
		}

		migrationsFileName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), "initial")
		migrationsPath := projectPath + common.BackendDirectory + common.MigrationsDirectory + migrationsFileName
		err = common.GenerateFile(migrationsPath, common.SQLFileExtension, migrationsFileConfiguration)
		if err != nil {
			panic(err)
		}

	}
	//if frontend == true {
	//	err := common.GenerateStubs(projectPath, common.FrontendDirectory)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

}
