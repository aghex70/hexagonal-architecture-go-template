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

	version, err := common.ScanStringWithDefault(common.ScanVersion, common.DefaultVersion)
	if err != nil {
		panic(err)
	}

	description, err := common.ScanStringWithDefault(common.ScanDescription, common.DefaultDescription)
	if err != nil {
		panic(err)
	}

	tc := templates.TemplateContext{
		Module:             projectModule,
		ProjectName:        projectName,
		ProjectDescription: description,
		ProjectVersion:     version,
	}

	cc := templates.ComposeConfiguration{}

	sequel, err := common.ScanStringCastBoolean(common.ScanSqlDatabase, common.DefaultActive)
	if err != nil {
		panic(err)
	}

	if sequel == true {
		mysql, err := common.ScanStringCastBoolean(common.ScanMysql, common.DefaultActive)
		if err != nil {
			panic(err)
		}

		if mysql == false {
			postgres, err := common.ScanStringCastBoolean(common.ScanPostgres, common.DefaultActive)
			if err != nil {
				panic(err)
			}
			cc.Postgres = postgres
		}
		cc.MySQL = mysql
	}

	if sequel == false {
		nosequel, err := common.ScanStringCastBoolean(common.ScanNoSqlDatabase, common.DefaultActive)
		if err != nil {
			panic(err)
		}

		if nosequel == true {
			mongodb, err := common.ScanStringCastBoolean(common.ScanMongodb, common.DefaultActive)
			if err != nil {
				panic(err)
			}
			cc.MongoDB = mongodb
		}
	}

	redis, err := common.ScanStringCastBoolean(common.ScanRedis, common.DefaultActive)
	if err != nil {
		panic(err)
	}
	cc.Redis = redis

	nginx, err := common.ScanStringCastBoolean(common.ScanReverseProxy, common.DefaultInactive)
	if err != nil {
		panic(err)
	}
	cc.NGINX = nginx

	rest, err := common.ScanStringCastBoolean(common.ScanRest, common.DefaultActive)
	if err != nil {
		panic(err)
	}

	grpc, err := common.ScanStringCastBoolean(common.ScanGrpc, common.DefaultActive)
	if err != nil {
		panic(err)
	}

	frontend, err := common.ScanStringCastBoolean(common.ScanReact, common.DefaultActive)
	if err != nil {
		panic(err)
	}

	////graphql, err := scanStringCastBoolean(ScanGraphql, DefaultInactive)
	//_, err = common.ScanStringCastBoolean(common.ScanGraphql, common.DefaultInactive)
	//if err != nil {
	//	panic(err)
	//}

	entities, err := common.ScanMultipleStrings(common.ScanEntities, common.DefaultEntity)
	if err != nil {
		panic(err)
	}

	projectPath := common.GenerateProjectPath(projectBasePath, projectName)
	err = common.CreateDirectory(projectPath)
	if err != nil {
		panic(err)
	}

	err = common.GenerateBackendStubs(projectPath)
	if err != nil {
		panic(err)
	}

	// Configuration
	configurationPath := projectPath + common.BackendDirectory + common.ConfigDirectory

	// Cache
	if redis == true {
		cacheFileConfiguration := templates.GetCacheFileConfiguration(tc)
		cacheConfigurationPath := configurationPath + common.CacheFileName
		err = common.GenerateFile(cacheConfigurationPath, common.GolangFileExtension, cacheFileConfiguration)
		if err != nil {
			panic(err)
		}
	}

	// Database
	if sequel == true {
		var databaseFileConfiguration []templates.FileConfiguration
		switch cc.MySQL {
		case true:
			databaseFileConfiguration = templates.GetMySQLFileConfiguration(tc)
		case false:
			//databaseFileConfiguration := templates.GetPostgresFileConfiguration(tc)
			databaseFileConfiguration = templates.GetMySQLFileConfiguration(tc)
		}
		databaseConfigurationPath := configurationPath + common.MySQLFileName
		err = common.GenerateFile(databaseConfigurationPath, common.GolangFileExtension, databaseFileConfiguration)
		if err != nil {
			panic(err)
		}
	}

	// gRPC
	if grpc == true {
		gRPCFileConfiguration := templates.GetGRPCFileConfiguration(tc)
		gRPCConfigurationPath := configurationPath + common.GRPCFileName
		err = common.GenerateFile(gRPCConfigurationPath, common.GolangFileExtension, gRPCFileConfiguration)
		if err != nil {
			panic(err)
		}
	}

	// REST
	if rest == true {
		restFileConfiguration := templates.GetRestFileConfiguration(tc)
		restConfigurationPath := configurationPath + common.RestFileName
		err = common.GenerateFile(restConfigurationPath, common.GolangFileExtension, restFileConfiguration)
		if err != nil {
			panic(err)
		}
	}

	// Server
	serverFileConfiguration := templates.GetServerFileConfiguration(tc)
	serverConfigurationPath := configurationPath + common.ServerFileName
	err = common.GenerateFile(serverConfigurationPath, common.GolangFileExtension, serverFileConfiguration)
	if err != nil {
		panic(err)
	}

	if len(entities) > 0 {
		for _, e := range entities {
			le := strings.ToLower(e)
			initial := string(le[0])
			tc.Entity = e
			tc.LowerEntity = le
			tc.Initial = initial

			// Domain
			domainPath := projectPath + common.BackendDirectory + common.DomainDirectory + le
			domainFileConfiguration := templates.GetDomainFileConfiguration(tc)
			err = common.GenerateFile(domainPath, common.GolangFileExtension, domainFileConfiguration)
			if err != nil {
				panic(err)
			}

			// Repositories
			if sequel == true {
				repositoryPath := projectPath + common.BackendDirectory + common.RepositoriesDirectory + le + "/"
				err = common.CreateDirectory(repositoryPath)
				if err != nil {
					panic(err)
				}

				repositoryFileConfiguration := templates.GetRepositoryFileConfiguration(tc)
				err = common.GenerateFile(repositoryPath+common.GormFileName, common.GolangFileExtension, repositoryFileConfiguration)
				if err != nil {
					panic(err)
				}
			}

			// Services
			servicePath := projectPath + common.BackendDirectory + common.BackendServicesDirectory + le + "/"
			err = common.CreateDirectory(servicePath)
			if err != nil {
				panic(err)
			}

			serviceFileConfigurations := templates.GetServiceFileConfiguration(tc)
			err = common.GenerateFile(servicePath+common.ServiceFileName, common.GolangFileExtension, serviceFileConfigurations)
			if err != nil {
				panic(err)
			}

			serviceTestFileConfiguration := templates.GetLowerEntityFileConfiguration(tc)
			err = common.GenerateFile(servicePath+common.ServiceTestFileName, common.GolangFileExtension, serviceTestFileConfiguration)
			if err != nil {
				panic(err)
			}

			validatorsFileConfiguration := templates.GetLowerEntityFileConfiguration(tc)
			err = common.GenerateFile(servicePath+common.ServiceValidatorsFileName, common.GolangFileExtension, validatorsFileConfiguration)
			if err != nil {
				panic(err)
			}

			// Handlers
			handlerPath := projectPath + common.BackendDirectory + common.HandlersDirectory + le + "/"
			err = common.CreateDirectory(handlerPath)
			if err != nil {
				panic(err)
			}

			restHandlerFileConfiguration := templates.GetRestHandlerFileConfiguration(tc)
			err = common.GenerateFile(handlerPath+common.RestHandlerFileName, common.GolangFileExtension, restHandlerFileConfiguration)
			if err != nil {
				panic(err)
			}

			restHandlerTestFileConfiguration := templates.GetLowerEntityFileConfiguration(tc)
			err = common.GenerateFile(handlerPath+common.RestHandlerTestFileName, common.GolangFileExtension, restHandlerTestFileConfiguration)
			if err != nil {
				panic(err)
			}

		}

		tc.Module = projectModule
		tc.ProjectName = projectName

		// Requests
		requestsFileConfigurations := templates.GetRequestsFileConfiguration(entities, tc)
		requestsPath := projectPath + common.BackendDirectory + common.PortsDirectory + common.RequestsFileName
		err = common.GenerateFile(requestsPath, common.GolangFileExtension, requestsFileConfigurations)
		if err != nil {
			panic(err)
		}

		// Services
		servicesFileConfigurations := templates.GetServicesFileConfiguration(entities, tc)
		servicesPath := projectPath + common.BackendDirectory + common.PortsDirectory + common.ServicesFileName
		err = common.GenerateFile(servicesPath, common.GolangFileExtension, servicesFileConfigurations)
		if err != nil {
			panic(err)
		}

		// Migrations
		migrationsFileConfigurations := templates.GetMigrationFileConfiguration(entities, tc)
		migrationsFileName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), "initial")
		migrationsPath := projectPath + common.BackendDirectory + common.MigrationsDirectory + migrationsFileName
		err = common.GenerateFile(migrationsPath, common.SQLFileExtension, migrationsFileConfigurations)
		if err != nil {
			panic(err)
		}

		// README.md
		readmeFileConfiguration := templates.GetReadmeFileConfiguration(tc)
		readmePath := projectPath + common.ReadmeFileName
		err = common.GenerateFile(readmePath, common.MarkdownFileExtension, readmeFileConfiguration)
		if err != nil {
			panic(err)
		}

		// commands
		serveFileConfigurations := templates.GetServeFileConfiguration(entities, tc)
		servePath := projectPath + common.BackendDirectory + common.CommandDirectory + common.ServeFileName
		err = common.GenerateFile(servePath, common.GolangFileExtension, serveFileConfigurations)
		if err != nil {
			panic(err)
		}

		makeMigrationsConfiguration := templates.GetMakeMigrationsFileConfiguration(tc)
		makeMigrationsPath := projectPath + common.BackendDirectory + common.CommandDirectory + common.MakeMigrationsFileName
		err = common.GenerateFile(makeMigrationsPath, common.GolangFileExtension, makeMigrationsConfiguration)
		if err != nil {
			panic(err)
		}

		migrateConfiguration := templates.GetMigrateFileConfiguration(tc)
		migratePath := projectPath + common.BackendDirectory + common.CommandDirectory + common.MigrateFileName
		err = common.GenerateFile(migratePath, common.GolangFileExtension, migrateConfiguration)
		if err != nil {
			panic(err)
		}

		rootFileConfiguration := templates.GetRootFileConfiguration(tc)
		rootPath := projectPath + common.BackendDirectory + common.CommandDirectory + common.RootFileName
		err = common.GenerateFile(rootPath, common.GolangFileExtension, rootFileConfiguration)
		if err != nil {
			panic(err)
		}

	}

	if frontend == true {
		cc.Frontend = true
		err := common.GenerateFrontendStubs(projectPath)
		if err != nil {
			panic(err)
		}
	}

	// docker-compose
	composeFileConfiguration := templates.GetDockerComposeFileConfiguration(cc, tc)
	composePath := projectPath + common.DockerComposeFileName
	err = common.GenerateFile(composePath, common.YAMLFileExtension, composeFileConfiguration)
	if err != nil {
		panic(err)
	}

}
