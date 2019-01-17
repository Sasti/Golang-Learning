# mysql-driver and sql package example
This project contains a small go programm that connects with a MySQL database. The Progremm executes several querys to change the state of the database and shows how to evaluate the results of a query.

The Programm only uses the basic sql package as an interface for the MySQL driver in use (github.com/go-sql-driver/mysql). There are several wrappers available to make the usage of the sql interface more convienient. For the sake of simplicity no wrapper has been used in this example.

For further reading on go sql package wrappers please take a look at the following projects

- https://github.com/avelino/awesome-go
- https://github.com/jmoiron/sqlx
- http://gorm.io

## Docker cotnainer for tests
For easier usage of the example you will find a docker-compose.yml in this directory. The compose file declares a mysql service and configures it to be available for the test programm. Feel free to star the container so that you don't need to setup your own mysql server.

    # Execute the docker-compose command in the directory of this README
    docker-compose up -d

## Start the go programm
You can compile and run the programm without additional tools. For simplicity just use the following command to compile and execute the programm with your locally install go environment.

    go run main.go