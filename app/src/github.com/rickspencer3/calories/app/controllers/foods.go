package controllers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
  "github.com/rickspencer3/calories/app/models"
	"log"
	"fmt"
)

type Foods struct {
	*revel.Controller
}

func (c Foods) List() revel.Result {
  fs := []models.Food{}

	db, err := sql.Open("mysql",
		"my_user:my_password@tcp(mariadb:3306)/calories")
	if err != nil {
		log.Fatal(err)
	}
	stm := "SELECT Description, id from foods;"
	rows, err := db.Query(stm)
  if err != nil {
		log.Fatal(err)
	}

  for rows.Next() {
    var id int
    var description string

    _ = rows.Scan(&description, &id)
		fmt.Println(description)
    fs = append(fs, models.Food{id,description})
  }
  //defer rows.Close()
	defer db.Close()

	return c.RenderJson(fs)
}
