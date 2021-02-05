# Go RESTful API POC (Usermanagement)


This POC is designed to get you up and running with a project structure optimized for developing
RESTful API services in Go. It promotes the best practices that follow the [SOLID principles](https://en.wikipedia.org/wiki/SOLID)
and [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). 
It encourages writing clean and idiomatic Go code. 

This provides the following features right out of the box:

* RESTful endpoints in the widely accepted format
* Standard CRUD operations of a database table
* Environment dependent application configuration management
* Structured logging with contextual information
* Error handling with proper error response generation
* Database migration
* Data validation


Gaps:
* JWT-based authentication
* Full test coverage
* Live reloading during development

 
The kit uses the following Go packages which can be easily replaced with your own favorite ones
since their usages are mostly localized and abstracted. 

* Web Framework: [gin-gonic](https://github.com/gin-gonic/gin)
* ORM: [go-gorm](https://github.com/go-gorm/gorm)
* Data validation: [go-validator](https://github.com/go-validator/validator/tree/v2)
* Logging: [zap](https://github.com/uber-go/zap)


## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires **Go 1.15 or above**.

[Docker](https://www.docker.com/get-started) is also needed if you want to try the kit without setting up your
own database server. The kit requires **Docker 17.05 or higher** for the multi-stage build support.

After installing Go and Docker, run the following commands to start experiencing this starter kit:

```shell
# download the POC
git clone https://github.com/jonayrodriguez/usermanagement.git

cd usermanagement
go mod vendor

# start a MySQL database server in a Docker container

docker run --name mysql8 -p3306:3306 -e MYSQL_ROOT_PASSWORD=pass4root -d mysql:8

**NOTE:** A schema called usermanagement needs to be created

# run the RESTful API server ( this will change to use make)
cd usermanagement/cmd/usermanagement
go run main.go 

```

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

* `POST api/v1/users`: Create a User
* `GET apit/v1/users/:username`: Get a user
* `GET api/v1/users`: Get all users (missing pagination)
* `DELETE api/v1/users/:username`: deletes a user

**TODO**
* `GET /healthcheck`: a healthcheck service provided for health checking purpose (needed when implementing a server cluster)
* `PATCH api/v1/users/:username`: Parcial user update
* `PUT apie/v1/users/:username`: Full user update


If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the following 
more complex scenarios:

```shell
# An example of creating a user: POST api/v1/users
curl -X POST -H "Content-Type: application/json" -d '{"username": "jonay", "surname": "rodriguez", "email":"jonay@gmail"}' http://localhost:8080/api/v1/users
# should return 201 and the user data in the response
```

## Project Layout

The starter kit uses the following project layout:
 
```
.
├── cmd                  main applications of the project
│   └── usermanagement   usermanagement main go file
├── config               configuration files for different environments
├── internal             private application and library code
│   ├── config           configuration library
│   ├── controller       domain logic structured (This should be renamed to service layer)
│   ├── database         configuration library
│   ├── log              logger definitions and context-aware logger
│   ├── router           API definitions (This could be renamed to api and being part of the /api folder)
├── migration            database migrations (TODO)
├── vendor               dependency management  (TODO)
├── development          deployment configurations and templates (TODO)
├── middleware           middlewares to use (TODO)
├── pkg                  public library code
    ├── model           entity/models definitions
    
```

**NOTE: Still missing some folder such as test, docs, etc. keep in mind that some refactorization needs to be done to have a full 3 layer (Api, service and repository)**


The top level directories `cmd`, `internal`, `pkg` are commonly found in other popular Go projects, as explained in
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Within `internal` and `pkg`, packages are structured by features in order to achieve the so-called
[screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html). For example, 
the `album` directory contains the application logic related with the album feature. 

Within each feature package, code are organized in layers (API, service, repository), following the dependency guidelines
as described in the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).


## Common Development Tasks

This section describes some common development tasks using this starter kit.

### Implementing a New Feature

**TODO**

### Working with DB Transactions

**TODO**


### Updating Database Schema

The starter kit uses [database migration](https://en.wikipedia.org/wiki/Schema_migration) to manage the changes of the 
database schema over the whole project development phase. The following commands are commonly used with regard to database
schema changes:

```shell
# Execute new migrations made by you or other team members.
# Usually you should run this command each time after you pull new code from the code repo. 
make migrate

# Create a new database migration.
# In the generated `migrations/*.up.sql` file, write the SQL statements that implement the schema changes.
# In the `*.down.sql` file, write the SQL statements that revert the schema changes.
make migrate-new

# Revert the last database migration.
# This is often used when a migration has some issues and needs to be reverted.
make migrate-down

# Clean up the database and rerun the migrations from the very beginning.
# Note that this command will first erase all data and tables in the database, and then
# run all migrations. 
make migrate-reset
```

### Managing Configurations

The application configuration is represented in `internal/config/config.go`. When the application starts,
it loads the configuration from a configuration file as well as environment variables. The path to the configuration 
file is specified via the `-config` command line argument which defaults to `./config/dev.yml`. Configurations
specified in environment variables should be named with the `APP_` prefix and in upper case. When a configuration
is specified in both a configuration file and an environment variable, the latter takes precedence. 

The `config` directory contains the configuration files named after different environments. For example,
`config/local.yml` corresponds to the local development environment and is used when running the application 
via `make run`.

Do not keep secrets in the configuration files. Provide them via environment variables instead. For example,
you should provide `Config.DSN` using the `APP_DSN` environment variable. Secrets can be populated from a secret
storage (e.g. HashiCorp Vault) into environment variables in a bootstrap script (e.g. `cmd/server/entryscript.sh`). 

## Deployment

The application can be run as a docker container. You can use `make build-docker` to build the application 
into a docker image. The docker container starts with the `cmd/server/entryscript.sh` script which reads 
the `APP_ENV` environment variable to determine which configuration file to use. For example,
if `APP_ENV` is `qa`, the application will be started with the `config/qa.yml` configuration file.

You can also run `make build` to build an executable binary named `main`. Then start the API server using the following
command,

```shell
./main -config=./config/prod.yml
```

```
