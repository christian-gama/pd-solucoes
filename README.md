## Passei Direto API

This is a RESTful API built using GoLang and PostgreSQL for the Passei Direto college. The API provides CRUD operations for the entity-relationship model of the college.

## Technologies Used

- GoLang
- PostgreSQL

## Requirements

- GoLang
- PostgreSQL
- Git
- Docker (optional)

## Getting Started

### Cloning the repository

```bash
git clone "https://github.com/christian-gama/pg-solucoes"
```

### Setting up the environment

The project needs three environment variables to run, they are:

- .env.dev
- .env.test
- .env.prod

The .env.dev file is used to run the project locally, the .env.test file is used to run the tests and the .env.prod file is used to run the project in production. There is a .env.example file that can be used as a template for the .env files.

### Initializing the project

To initialize the project, run the following command:

```bash
make init
```

### Running the tests

To run the tests, run the following command:

```bash
make test
```

It will start a PostgreSQL container, run migrations and run the tests.
