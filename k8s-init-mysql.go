package main

import (
		"github.com/go-sql-driver/mysql"
		"database/sql"
    "fmt"
    "time"
    "github.com/namsral/flag"
)

func main() {

  userPtr := flag.String("dbuser", "", "User to connect as (Required)")
  passPtr := flag.String("dbpassword", "", "Password to connect with (Required)")
  addrPtr := flag.String("dbhost", "", "Host IP or DNS connect to - TCP Only now (Required)")
  dbPtr := flag.String("dbname", "", "Database name (Required)")
  timeoutPtr := flag.String("timeout", "", "Database TCP timeout (OS default)")
  sqlPtr := flag.String("dbsql", "select 1;", "SQL to test (\"select 1;\")")
  flag.Parse()

  config := mysql.NewConfig()
  config.User =  *userPtr
  config.Passwd = *passPtr
  config.Net = "tcp"
  config.Addr = *addrPtr
  config.DBName = *dbPtr

  if *timeoutPtr != "" {
    parsedTimeout, err := time.ParseDuration(*timeoutPtr)
    config.Timeout = parsedTimeout
    checkErr(err)
  }

  dsn := config.FormatDSN()
  fmt.Println(dsn)
  db, err := sql.Open("mysql", dsn)
  checkErr(err)
  defer db.Close()

  err = db.Ping()
  checkErr(err)

  stmtOut, err := db.Prepare(*sqlPtr)
  checkErr(err)
  defer stmtOut.Close() 

  fmt.Println("Database connection established")

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
