# This is part 1 of a series of courses from [here](https://www.udemy.com/course/htmx-go-build-fullstack-applications-with-golang-and-htmx/)

This part is sort of an introduction to some concepts.

- Here we will use `Docker` and `MySQL`

- MySQL - `docker run --name demo-mysql -e MYSQL_ROOT_PASSWORD=root -p 3333:3306 -d mysql:latest`
- Inside the container - `mysql -u root -p`
- Check for DB:
    `SELECT * FROM user;`
- If nothing showing up - `use mysql;` and repeat the one above.

# Base Project Creation: - [Base Project](./go-htmx-app/)

- init project : `go mod init github.com/BugzTheBunny/go-htmx-app`
- installing gorilla mux : `go get -u github.com/gorilla/mux`
- Installing MySQL driver : `go get -u github.com/go-sql-driver/mysql`
