package main

import (
	"github.com/aghex70/hexagonal-architecture-go-template/common"
	"github.com/aghex70/hexagonal-architecture-go-template/templates"
	"os"
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
		var es []templates.DomainData
		for _, e := range entities {
			entity := templates.DomainData{
				Entity: e,
				Module: projectModule,
			}
			err = common.GenerateFile(projectPath, common.BackendDirectory, common.DomainDirectory, e, templates.DomainTemplate, entity)
			if err != nil {
				panic(err)
			}
			es = append(es, entity)
		}

		err = common.GenerateMultiPartsFile(projectPath, common.BackendDirectory, common.PortsDirectory, common.RequestsFileName, templates.RequestsStartTemplate, templates.RequestsEntitiesTemplate, es)
		if err != nil {
			panic(err)
		}

		err = common.GenerateMultiPartsFile(projectPath, common.BackendDirectory, common.PortsDirectory, common.ServicesFileName, templates.ServicesStartTemplate, templates.ServicesEntitiesTemplate, es)
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
