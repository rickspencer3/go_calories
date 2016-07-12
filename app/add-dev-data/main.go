package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
  "bytes"
	"log"
  "io"
  "os"
	"strings"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql",
		"my_user:my_password@tcp(mariadb:3306)/calories")
	if err != nil {
		log.Fatal(err)
	} else {
		executeString(db, readFile("calorielookup.sql"))
		foodsString := readFile("foods.sql")
		foodsStatements := strings.Split(foodsString, ";\n")
		for _ ,s := range foodsStatements {
			executeString(db, s)
		}
	}
	defer db.Close()
}

func executeString(db *sql.DB, statement string) {
	_, err := db.Exec(statement)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(".")
}

func readFile(filename string) string {
  s := ""
  f , err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  } else {
    buf := bytes.NewBuffer(nil)
    io.Copy(buf, f)
    s = string(buf.Bytes())
    f.Close()
  }
  return s
}
