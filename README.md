## Book_store
This is a Go-based REST API for a Book Store. It uses mux for routing, GORM as an ORM library, and PostgreSQL as the database. Follow the instructions below to set up and run the project.

Requirements
To run this project, you need the following:

- [GO 1.23.0](https://go.dev/dl/) 

- [PostgreSQL](https://www.postgresql.org/download/)

- Git

## Setup
1. Clone the Repository

```sh
git clone https://github.com/moheddine-belhaj/Book_store.git
cd Book_store

```

2. Set Up the PostgreSQL Database
Create a PostgreSQL database and note the connection details (host, port, user, password, dbname) as they are required for configuring the environment.

3. Install Dependencies

#### Routing
Install the `Gorilla mux` router for handling routes:

```sh
go get github.com/gorilla/mux
```

#### Install GORM and the PostgreSQL driver:

```sh 
go get gorm.io/gorm

```
```sh
go get gorm.io/driver/postgres
```
#### Install the cors package for Cross-Origin Resource Sharing support:

```sh
go get github.com/rs/cors
```

#### Install additional PostgreSQL-related packages:

```sh
go get github.com/jackc/pgx/v5
go get github.com/jinzhu/inflection
go get github.com/jinzhu/now

```

## Install Dependencies

```sh
go mod tidy
```

## Run the Application

```sh
go run main.go
```