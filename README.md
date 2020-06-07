![headline](images/logo.jpeg)

# [GO-GENERATOR](https://github.com/wilian746/go-generator)

<p align="center">
  <a href="https://github.com/wilian746/go-generator/actions"><img src="https://img.shields.io/github/workflow/status/wilian746/go-generator/Go/master?label=Build"/></a>
  <a href="https://github.com/wilian746/go-generator/releases"><img src="https://img.shields.io/github/v/tag/wilian746/go-generator?color=green&label=Version"/></a>
</p>

## What is GO-GENERATOR?
Go generator is an simple command line interface to generate default files to start your implementation with go in a base pre-configurated.
Currently an generator of code is:
- [GORM](https://google.com) - This standart project integrated with relational database your connections is:
    - PostgreSQL
    - MySQL
    - SQLServer
    - SQLite3

## Usage
### Install in your local machine
```bash
go get -u github.com/wilian746/go-generator/cmd/go-generator
```

### Check installation
```bash
go-generator
```

### Init application
This command will copy all standart content using gorm library to path and module indicated  
 ```bash
go-generator init gorm app
 ```
After running the command above it will ask you which is the directory you want to perform the standard installation.
By default, it's already suggests the current directory as the installation location, but you can change it.
See example!
```bash
✔ Enter the full path of the directory destiny!: /home/wilian/go/src/github.com/wilian746/go-generator/tmp
```
After informing the installation destination it will ask what is the name of the module you want to add to your application.
By default, it's already suggests this project module name, but you can change it too.
See example!
```bash
✔ Enter module of golang project: github.com/wilian746/go-generator/tmp
```
🤩 Yeaahhh!! Your installation's finished! 😁
    

## Generated structure
### standart-gorm
This project follows the standard structure of the [golang-standard](https://github.com/golang-standards/project-layout).
We generatated some codes to facilities the life of the developer.
- `/cmd` This folder contains the main project, you can see that the integration with the database starts here and it is passed on to the rest of the project. With that we solved some problems with the loss of connection with the database;
- `/config` This folder contains the environments of the project they are:
    - `PORT` -> Port the application will run on;
    - `TIMEOUT` -> Timeout of routes in minutes;
    - `DATABASE_DIALECT` -> [Dialect usage in GORM](https://gorm.io/docs/connecting_to_the_database.html);
    - `DATABASE_URI` -> [Connection string usage in GORM](https://gorm.io/docs/connecting_to_the_database.html);
- `/deploymets` This folder contains the files of deployments to run dependences of server; 
- `/docs` This folder contains the docs of project using Swagger, you can this example of how to update/create using [swag](https://github.com/swaggo/swag); 
- `/migrations` This folder contains the migrations files with drivers specifics; 
    - `/migrations/{DRIVER}` The folders are driver with your migrations files to up or down versions; 
- `/internal` This folder contains the internal implementations;
    - `/internal/controllers` This folder contains internal rules and conditions treatments of your application;
    - `/internal/entities` This folder contains entities usage to save in your database and manipulate data in project;
    - `/internal/handlers` This folder contains input and output response HTTP. By default, we implement two routes for you `/health` to check if connection with a database is alive, and a CRUD of products in route `/product`;
    - `/internal/routes` This folder contains the settings of the routes, middleware, implementation of the routes;
    - `/internal/rules` This folder contains the rules of entities they are: required fields, conversions, etc;
    - `/internal/utils` This folder contains all the methods that can be used throughout the development of the project and reused as calls for environment variables, standardization of HTTP responses and standardization of the displayed logs;
- `/pkg/repository` This folder contains connection with a database;
    - `/pkg/repository/adapter` This folder contains the Adapter usage to manipulate a query's in the database;
    - `/pkg/repository/database` This folder contains the Connection with database and validation of inputs;
    - `/pkg/repository/entities` This folder contains the Interface recommend to usage in your entities;
    - `/pkg/repository/response` This folder contains the Response your objective is all there queries executed will be standardized for a standard response, thus facilitating error handling and data conversions if necessary.

## Plans?who generated
We are just getting started, the main objective is to aggregate several banks on a solid and consistent implementation basis so that all developers can save development time, so we have some activities that we will still do in the short, medium and long term.
* Phase 1: Initial project implementation using relational database ✔ 
* Phase 2: Integration with the MongoDB database ⌚️
* Phase 3: Integration with the RabbitMQ broker💡 ️
* Phase 4: Creation of separate resources such as creating only a new route or a new controller ✴️
* Phase 5: Add optional implementations like the Redis memory database in specific routes ♨️
- And several other ideas...👁‍ 

## Issue?
We are happy with your help, you can direct us by channel of [Issues](https://github.com/wilian746/go-generator/issues) that we will help with the greatest pleasure!

## Contributing
Nice! Welcome to the team then! Just make your modification and open the [pull request](https://github.com/wilian746/go-generator/pulls). We ask that the branch of [develop](https://github.com/wilian746/go-generator/tree/develop) is always placed as a destination, because that way we can move up to production safe and tested implementations.


# Thank's for usage ! ✌️