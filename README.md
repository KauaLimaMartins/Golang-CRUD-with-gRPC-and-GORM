# Golang CRUD with gRPC server and Gorm

A simple golang CRUD using gRPC server and gorm

## How to run it

First, you need to have [Docker](https://www.docker.com/) installed

- Clone the project:
```bash
git clone https://github.com/KauaLimaMartins/Golang-CRUD-with-gRPC-and-GORM.git
```

- Open it in your terminal and run:
```bash
docker network create -d bridge my-bridge
```

- And now, run:
```bash
docker compose up
```

- Now, you can access it by calling a gRPC request in localhost:8000 like this:

![Insomnia gRPC request example](https://github.com/KauaLimaMartins/Golang-CRUD-with-gRPC-and-GORM/blob/master/github/readme-assets/create.png)

![Insomnia gRPC request example](https://github.com/KauaLimaMartins/Golang-CRUD-with-gRPC-and-GORM/blob/master/github/readme-assets/get-all.png)
