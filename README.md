# Kubernetes Init Container Helper - MySQL Probe

This is a simple GO application that can be staticlly compiled and build into
a Stratch docker container and used as a K8S init contianer to check that 
a MySQL service is alive and usable with provided credentials and optionally
running a SQL Query that will result in a non-empty response

## Getting Started

Clone down and run 
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```

### Prerequisites

		"github.com/go-sql-driver/mysql"
    "github.com/namsral/flag"

```
		go get github.com/go-sql-driver/mysql
    go get github.com/namsral/flag
```

## Usage ##

### Command line arguments ###
Usage of /k8s-init-mysql:
  -dbhost string
    	Host IP or DNS connect to - TCP Only now (Required)
  -dbname string
    	Database name (Required)
  -dbpassword string
    	Password to connect with (Required)
  -dbsql string
    	SQL to test ("select 1;") (default "select 1;")
  -dbuser string
    	User to connect as (Required)

### Environment Variables ###

DBHOST
DBNAME
DBPASSWORD
DBSQL
DBUSER

The SQL is passed in as is - you will need to escape it yourself.

### Example Usage ###

```
docker run -t -e DBHOST=db.default.svc.cluster.local -e DBUSER=someone \
            -e DBPASSWORD=somepass -e DBNAME=somedb \
            craftypenguins/k8s-init-mysql:latest 
```

## Authors

* **Richard Clark** - *Initial work* - [kti-richard](https://github.com/kti-richard)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

