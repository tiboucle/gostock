# gostock
Inventory REST API built in Go and SQLite

### Prerequisite
1. [Go](https://golang.org/)
```sh
$ go version
go version go1.10.3 linux/amd64
```
2. [SQLite](https://www.sqlite.org/index.html)
```sh
$ sqlite3 version
SQLite version 3.22.0 2018-01-22 18:45:57
```

### How to Run
1. Clone this repository
```sh
$ git clone https://github.com/tiboucle/gostock.git && cd goventory
```
2. Get additional library
```sh
$ go get github.com/jehiah/go-strftime github.com/mattn/go-sqlite3
```
3. Execute `run.sh`
```sh
$ ./run.sh
```

3. `gostock` listening on port 8090
```sh
Starting web server GoStock on http://localhost:8090
```

### API Test
1. Items Stock:                 `http://localhost:8090/api/items`

2. Items Stock Incoming:        `http://localhost:8090/api/incoming/items`

3. Items Stock Outgoing/Sales:  `http://localhost:8090/api/sales/items`

4. Sales Item Value Report:     `http://localhost:8090/api/sales/item/value/report`

5. Sales Report:                `http://localhost:8090/api/sales/item/report?startdate=2018-01-07&enddate=2018-01-08`

Both `startdate` and `enddate` requires date in `YYYY-MM-DD` format.

Either way, it throws `400 Bad Request`.
