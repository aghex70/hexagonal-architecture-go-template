# hexagonal-architecture-go-template

A command line utility which aims to create an hexagonal architecture templater for Golang projects

## Features
 - Dockerized architecture generation.
 - Backend generation with different layers: domain, services & repositories.
 - React frontend generation (WIP)
 - Layers interface generation.
 - Basic CRUD operations implemented for handlers, services & repositories.
 - MySQL support.
 - PostgreSQL support (WIP)
 
### Usage

- Simple command line usage:

  ```bash
  # Compile the code
  $ go build main.go 
  # Run the templater providing the desired path where the new project is going to be created
  $ go run main.go /<DESIRED_FOLDER>/
  ```
  
- User will be prompted for input to determine and build the new project's configuration and structure
  ```bash
  Project module (cannot be empty): github.com/aghex70/fake-project
  Project name (cannot be empty): fake-project
  Version [0.1.0]: 
  Description: Example generated project
  Add SQL database [Y/n]: 
  Add MySQL [Y/n]: 
  Add Redis [Y/n]: 
  Add reverse proxy (NGINX) [Y/n]: 
  Add REST API [Y/n]: 
  Add gRPC API [Y/n]: 
  Generate frontend (React) skeletons [Y/n]:  
  Entities [User]: Card User Purchase
  ```

### Generated architecture
 ```bash
.
├── backend
│   ├── cmd
│   │   ├── makemigrations.go
│   │   ├── migrate.go
│   │   ├── root.go
│   │   └── serve.go
│   ├── config
│   │   ├── cache.go
│   │   ├── config.go
│   │   ├── database.go
│   │   ├── grpc.go
│   │   ├── rest.go
│   │   └── server.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── internal
│   │   ├── core
│   │   │   ├── domain
│   │   │   │   ├── card.go
│   │   │   │   ├── purchase.go
│   │   │   │   └── user.go
│   │   │   ├── ports
│   │   │   │   ├── repositories.go
│   │   │   │   ├── requests.go
│   │   │   │   └── services.go
│   │   │   └── services
│   │   │       ├── card
│   │   │       │   ├── service.go
│   │   │       │   ├── service_test.go
│   │   │       │   └── validators.go
│   │   │       ├── purchase
│   │   │       │   ├── service.go
│   │   │       │   ├── service_test.go
│   │   │       │   └── validators.go
│   │   │       └── user
│   │   │           ├── service.go
│   │   │           ├── service_test.go
│   │   │           └── validators.go
│   │   ├── handlers
│   │   │   ├── card
│   │   │   │   ├── rest.go
│   │   │   │   └── rest_test.go
│   │   │   ├── purchase
│   │   │   │   ├── rest.go
│   │   │   │   └── rest_test.go
│   │   │   └── user
│   │   │       ├── rest.go
│   │   │       └── rest_test.go
│   │   └── stores
│   │       ├── card
│   │       │   └── gorm.go
│   │       ├── purchase
│   │       │   └── gorm.go
│   │       └── user
│   │           └── gorm.go
│   ├── main.go
│   ├── persistence
│   │   └── database
│   │       ├── database.go
│   │       ├── gorm.go
│   │       ├── migration.go
│   │       └── migrations
│   │           └── 20230302071933_initial.sql
│   └── server
│       ├── grpc.go
│       ├── rest.go
│       └── server.go
├── docker-compose.yml
├── frontend
│   └── src
│       ├── components
│       ├── feature-flags
│       ├── routes
│       ├── services
│       ├── use-cases
│       └── utils
└── README.md
  ```
