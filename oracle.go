package main

import (
	"database/sql"
	"log"

	"gopkg.in/gorp.v1"
	_ "gopkg.in/rana/ora.v3"
)

type dEPARTMENT struct {
	Department_name string
}

func main() {

	db, err := sql.Open("ora", "system/oracle@localhost:1521/xe")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	dbmapOracle := &gorp.DbMap{Db: db}

	var DEPARTMENTS []dEPARTMENT
	_, err = dbmapOracle.Select(&DEPARTMENTS, `select DEPARTMENT_NAME from HR.DEPARTMENTS`)
	if err != nil {
		log.Println(err)
	}

	for x, i := range DEPARTMENTS {
		log.Printf("    %d: %v\n", x, i)
	}

}
