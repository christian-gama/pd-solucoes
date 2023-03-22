## Passei Direto API

This is a RESTful API built using GoLang and PostgreSQL for the Passei Direto college. The API provides CRUD operations for the entity-relationship model of the college.
This project is following the principles of Clean Architecture, SOLID and Domain-Driven Design.

## Technologies Used

- GoLang
- PostgreSQL

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

### Setting up the environment

The project needs three environment variables to run, they are:

- .env.dev
- .env.test
- .env.prod

The .env.dev file is used to run the project locally, the .env.test file is used to run the tests and the .env.prod file is used to run the project in production. There is a .env.example file that can be used as a template for the .env files.

### Initializing the project

To initialize the project, run the following command (requires Docker):

```bash
make init
```

If you don't have Docker installed, you can run the commands inside the Makefile manually.

### Running the tests

To run the tests, run the following command:

```bash
make test
```

It will start a PostgreSQL container, run migrations and run the tests.
