package main

import (
	"fmt"
	"github.com/aghex70/hexagonal-architecture-go-template/common"
	"github.com/aghex70/hexagonal-architecture-go-template/templates"
	"os"
	"sort"
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

	ec := templates.ExtraConfiguration{}

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
			ec.Postgres = postgres
		}
		ec.MySQL = mysql
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
			ec.MongoDB = mongodb
		}
	}

	redis, err := common.ScanStringCastBoolean(common.ScanRedis, common.DefaultActive)
	if err != nil {
		panic(err)
	}
	ec.Redis = redis

	nginx, err := common.ScanStringCastBoolean(common.ScanReverseProxy, common.DefaultInactive)
	if err != nil {
		panic(err)
	}
	ec.NGINX = nginx

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
	sort.Strings(entities)
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

	configurationPath := projectPath + common.BackendDirectory + common.ConfigDirectory
	// Configuration
	configurationFileConfiguration := templates.GetConfigFileConfiguration(tc)
	configurationFilePath := configurationPath + common.ConfigurationFileName
	err = common.GenerateFile(configurationFilePath, common.GolangFileExtension, configurationFileConfiguration)
	if err != nil {
		panic(err)
	}

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
		switch ec.MySQL {
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
			tc.TableSuffix = templates.GenerateTableSuffix(tc.LowerEntity)

			// Domain
			domainPath := projectPath + common.BackendDirectory + common.DomainDirectory + le
			domainFileConfiguration := templates.GetDomainFileConfiguration(tc)
			err = common.GenerateFile(domainPath, common.GolangFileExtension, domainFileConfiguration)
			if err != nil {
				panic(err)
			}

			// gorm repositories
			if sequel == true {
				repositoryPath := projectPath + common.BackendDirectory + common.StoresDirectory + le + "/"
				err = common.CreateDirectory(repositoryPath)
				if err != nil {
					panic(err)
				}

				repositoryFileConfiguration := templates.GetGormRepositoryFileConfiguration(tc)
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

		// repository
		repositoryFileConfigurations := templates.GetRepositoryFileConfiguration(entities, tc)
		repositoryPath := projectPath + common.BackendDirectory + common.StoresDirectory + common.RepositoriesFileName
		err = common.GenerateFile(repositoryPath, common.GolangFileExtension, repositoryFileConfigurations)
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
		ec.Frontend = true
		err := common.GenerateFrontendStubs(projectPath)
		if err != nil {
			panic(err)
		}

		//	frontend env
		frontendEnvFileConfiguration := templates.GetEmptyFileConfiguration(tc)
		frontendEnvFilePath := projectPath + common.FrontendDirectory + common.FrontendEnvFileName
		err = common.GenerateFile(frontendEnvFilePath, "", frontendEnvFileConfiguration)
		if err != nil {
			panic(err)
		}
	}

	// docker-compose
	composeFileConfiguration := templates.GetDockerComposeFileConfiguration(ec, tc)
	composePath := projectPath + common.DockerComposeFileName
	err = common.GenerateFile(composePath, common.YAMLFileExtension, composeFileConfiguration)
	if err != nil {
		panic(err)
	}

	// main
	mainFileConfiguration := templates.GetMainFileConfiguration(tc)
	mainPath := projectPath + common.BackendDirectory + common.MainFileName
	err = common.GenerateFile(mainPath, common.GolangFileExtension, mainFileConfiguration)
	if err != nil {
		panic(err)
	}

	// go.mod
	goModFileConfiguration := templates.GetGoModFileConfiguration(tc)
	goModPath := projectPath + common.BackendDirectory + common.GoModFileName
	err = common.GenerateFile(goModPath, common.ModFileExtension, goModFileConfiguration)
	if err != nil {
		panic(err)
	}

	// backendDockerfile
	backendDockerFileConfiguration := templates.GetBackendDockerfileFileConfiguration(tc)
	backendDockerFilePath := projectPath + common.BackendDirectory + common.DockerFileName
	err = common.GenerateFile(backendDockerFilePath, "", backendDockerFileConfiguration)
	if err != nil {
		panic(err)
	}

	// persistence
	databaseFileConfiguration := templates.GetMySQLPersistenceFileConfiguration(tc)
	databaseFilePath := projectPath + common.BackendDirectory + common.DatabaseDirectory + common.DatabaseFileName
	err = common.GenerateFile(databaseFilePath, common.GolangFileExtension, databaseFileConfiguration)
	if err != nil {
		panic(err)
	}

	gormFileConfiguration := templates.GetGormPersistenceFileConfiguration(tc)
	gormFilePath := projectPath + common.BackendDirectory + common.DatabaseDirectory + common.GormFileName
	err = common.GenerateFile(gormFilePath, common.GolangFileExtension, gormFileConfiguration)
	if err != nil {
		panic(err)
	}

	migrationFileConfiguration := templates.GetMigrationPersistenceFileConfiguration(tc)
	migrationFilePath := projectPath + common.BackendDirectory + common.DatabaseDirectory + common.MigrationFileName
	err = common.GenerateFile(migrationFilePath, common.GolangFileExtension, migrationFileConfiguration)
	if err != nil {
		panic(err)
	}

	// server
	serverInterfaceFileConfiguration := templates.GetServerInterfaceFileConfiguration(tc)
	serverInterfaceFilePath := projectPath + common.BackendDirectory + common.ServerDirectory + common.ServerFileName
	err = common.GenerateFile(serverInterfaceFilePath, common.GolangFileExtension, serverInterfaceFileConfiguration)
	if err != nil {
		panic(err)
	}

	grpcFileConfiguration := templates.GetGRPCInterfaceFileConfiguration(tc)
	grpcFilePath := projectPath + common.BackendDirectory + common.ServerDirectory + common.GRPCFileName
	err = common.GenerateFile(grpcFilePath, common.GolangFileExtension, grpcFileConfiguration)
	if err != nil {
		panic(err)
	}

	restFileConfiguration := templates.GetRestServerFileConfiguration(entities, tc)
	restFilePath := projectPath + common.BackendDirectory + common.ServerDirectory + common.RestFileName
	err = common.GenerateFile(restFilePath, common.GolangFileExtension, restFileConfiguration)
	if err != nil {
		panic(err)
	}

	//	backend env
	backendEnvFileConfiguration := templates.GetBackendEnvFileConfiguration(ec, tc)
	backendEnvFilePath := projectPath + common.BackendDirectory + common.BackendEnvFileName
	err = common.GenerateFile(backendEnvFilePath, "", backendEnvFileConfiguration)
	if err != nil {
		panic(err)
	}

}
