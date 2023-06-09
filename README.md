## Passei Direto API

This is a RESTful API built using GoLang and PostgreSQL for the Passei Direto college. The API provides CRUD operations for the entity-relationship model of the college.
This project is following the principles of Clean Architecture, SOLID and Domain-Driven Design.

[API documentation](https://www.postman.com/christiangama/workspace/passei-direto-api/api/b77c83a2-2f9e-44ff-b453-55cf48af891f) available on Postman.

## Features
- All **GET** endpoints have **sorting**, **filtering** and **pagination** (consult API documentation for more details)
- All endpoints have validation for the request body
- Over 600 tests, including unit and integration tests
- All endpoints have documentation
- [Domain-Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design) principles
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) principles
- [SOLID](https://en.wikipedia.org/wiki/SOLID) principles
- CI/CD pipeline with GitHub Actions
- Docker containers for the API and the database
- Easy migrations using task automation
- Code linting
- Easy project setup by using Docker and Makefiles
- Git commit messages following the Conventional Commits specification
- RESTful API principles
- API versioning
- Middleware for logging and recovering from panics
- User friendly error messages

## Requirements

- GoLang
- PostgreSQL
- Git
- Docker (optional)
- Make (optional)

Even though Docker is optional, it is recommended to use it to run the project, as it will make the process of setting up the environment, dependencies and migrations easier. You can find the installation instructions for Docker [here](https://docs.docker.com/get-docker/).
If you plan to not use Docker, you will need to install the dependencies and run the migrations manually.

Make comes pre-installed in most Linux distributions, but if you are using Windows, you will need to install it manually. You can find the installation instructions [here](https://www.gnu.org/software/make/). It will simplify the process of running the project. If you don't plan to install Make, you can run the commands inside the Makefile manually.

## Getting Started

### Cloning the repository

```bash
git clone "https://github.com/christian-gama/pd-solucoes"
```

### Initializing the project

To initialize the project, run the following command (requires Docker):

```bash
make init
```
Optionally, you can also run the migration with seed, which will populate the database with some data (for development mode only), using the following command:

```bash
make init SEED=true
```

If you don't have Docker installed, you can run the commands inside the Makefile manually.

### Setting up the environment

The project needs three environment variables to run, they are:

- .env.dev
- .env.test
- .env.prod

If you ran the init command, the .env files will already be created, otherwise you can use the .env.example file as a template.

##### Please make sure the ports are available in your machine. If you find any issue, it's probably because of the ports.

### Running the tests

To run the tests, run the following command:

```bash
make test
```

It will start a PostgreSQL container, run migrations and run the tests.

### Running the project
##### Migration
If you ran the `make init` command, the migrations will already be run for development environment. Production environment wont run migrations automatically, so you will need to run them as described below:

```bash
make migrate-up ENV_FILE=.env.prod
```
Change from `prod` to `dev` if you want to run the migrations in the development environment. This instruction serves for the other commands that requires the environment variable.

##### Run project
Finally, you can run the project with the following command:

```bash
make docker-dev
```
The development mode will run with Gin, which can automatically reload the project when a file changes, making the development process easier.

**or**
```bash
make docker-prod
```
The production mode will build the project and run it with the binary file.
&nbsp;

##### For Linux users
If you are running the project in Linux, you may face permission issues because of the way Docker works with volumes. To solve this, you can run the command that failed with `sudo`, or you can change the permissions of the folder.
