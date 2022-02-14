# URL shortener API server

2 APIs to transform original URL into short URL and to find original URL by short URL

## Usage

* Build config.yml with config.example.yml

* Run DB migration

* Run main.go to start the server
```sh
$ go run main.go
```

### DB migration

* Install goose: https://github.com/pressly/goose

* Create a database named `Url_Shortener`

* Run goose up to latest migration (replace 'username' and 'password' with your own account)
```sh
$ goose -dir ./goose mysql "username:password@tcp(localhost:3306)/Url_Shortener?charset=utf8mb4&parseTime=True" up
```

* Reset database
```sh
$ goose -dir ./goose mysql "username:password@tcp(localhost:3306)/Url_Shortener?charset=utf8mb4&parseTime=True" down-to 0
```

* Create new migration file (replace 'test' to your version name)
```sh
$ goose -dir ./goose -s create test sql
```

### Optimization list

* Docker
* Unit test
* Problem of insufficient hash in the future
* Race condition problem