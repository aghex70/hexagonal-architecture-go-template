package hexagonal_architecture_go_template

func main() {
	projectName, err := scanString(SCAN_PROJECT_NAME)
	if err != nil {
		panic(err)
	}

	version, err := scanStringWithDefault(SCAN_VERSION, DEFAULT_VERSION)
	if err != nil {
		panic(err)
	}

	description, err := scanStringWithDefault(SCAN_DESCRIPTION, DEFAULT_DESCRIPTION)
	if err != nil {
		panic(err)
	}

	frontend, err := scanStringCastBoolean(SCAN_REACT, DEFAULT_ACTIVE)
	if err != nil {
		panic(err)
	}

	sequel, err := scanStringCastBoolean(SCAN_SQL_DATABASE, DEFAULT_ACTIVE)
	if err != nil {
		panic(err)
	}

	mysql, err := scanStringCastBoolean(SCAN_MYSQL, DEFAULT_INACTIVE)
	if err != nil {
		panic(err)
	}

	postgres, err := scanStringCastBoolean(SCAN_POSTGRES, DEFAULT_ACTIVE)
	if err != nil {
		panic(err)
	}

	nosequel, err := scanStringCastBoolean(SCAN_NO_SQL_DATABASE, DEFAULT_ACTIVE)
	if err != nil {
		panic(err)
	}

	mongodb, err := scanStringCastBoolean(SCAN_MONGODB, DEFAULT_ACTIVE)
	if err != nil {
		panic(err)
	}

	redis, err := scanStringCastBoolean(SCAN_REDIS, DEFAULT_ACTIVE)
	if err != nil {
		panic(err)
	}

	nginx, err := scanStringCastBoolean(SCAN_REVERSE_PROXY, DEFAULT_INACTIVE)
	if err != nil {
		panic(err)
	}

	api, err := scanStringCastBoolean(SCAN_API, DEFAULT_ACTIVE)
	if err != nil {
		panic(err)
	}

	rest, err := scanStringCastBoolean(SCAN_REST, DEFAULT_ACTIVE)
	if err != nil {
		panic(err)
	}

	grpc, err := scanStringCastBoolean(SCAN_GRPC, DEFAULT_ACTIVE)
	if err != nil {
		panic(err)
	}

	graphql, err := scanStringCastBoolean(SCAN_GRAPHQL, DEFAULT_INACTIVE)
	if err != nil {
		panic(err)
	}

	entities, err := scanMultipleStrings(SCAN_ENTITIES)
	if err != nil {
		panic(err)
	}

}
