## Passei Direto API

This is a RESTful API built using GoLang and PostgreSQL for the Passei Direto college. The API provides CRUD operations for the entity-relationship model of the college.
This project is following the principles of Clean Architecture, SOLID and Domain-Driven Design.

## Features
- All endpoints have sorting, filtering and pagination (consult API documentation for more details)
- All endpoints have validation for the request body
- Over 600 tests, including unit and integration tests
- All endpoints have documentation
- API documentation is available in Postman
- Domain-Driven Design and Clean Architecture principles
- SOLID principles
- CI/CD pipeline with GitHub Actions
- Docker containers for the API and the database
- Automatic migrations
- Code linting
- Easy project setup by using Docker and Makefiles
- Git commit messages following the Conventional Commits specification
- RESTful API principles
- API versioning
- Middleware for logging and recovering from panics

## Requirements

- GoLang
- PostgreSQL
- Git
- Docker (optional)

Even though Docker is optional, it is recommended to use it to run the project, as it will make the process of setting up the environment, dependencies and migrations easier. You can find the installation instructions for Docker [here](https://docs.docker.com/get-docker/).
If you plan to not use Docker, you will need to install the dependencies and run the migrations manually.

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

If you don't have Docker installed, you can run the commands inside the Makefile manually.

### Setting up the environment

The project needs three environment variables to run, they are:

- .env.dev
- .env.test
- .env.prod

If you ran the init command, the .env files will already be created, otherwise you can use the .env.example file as a template.

### Running the tests

To run the tests, run the following command:

```bash
make test
```

It will start a PostgreSQL container, run migrations and run the tests.

### Running the project
To run the project, run the following command:

```bash
make docker-dev
```
or
```bash
make docker-prod
```