# KO

A small GraphQL API used to create and play games of Go.

### Running the server

1. Start mysql;

```
$ docker run --name ko-mysql \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=example \
    -e MYSQL_DATABASE=ko \
    -d mysql
```

2. Run the app

```
$ go run main.go
```

### Run with Docker

TODO link mysql

```
$ docker build -t ko-app .
$ docker run -it -p 8080:8080 --rm ko-app
```

### Run feature tests

Run with Docker;

```
$ ./features/run.sh
```

Run without Docker (requires mysql to be installed with eg. `brew install mysql`);

```
$ pushd ./features; bundle install; popd
$ cucumber
```