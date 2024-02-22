package database

import (
	"cloudbees/adapters/structures"
	"cloudbees/adapters/util"
	sql "database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DBConnect() {
	var err error
	DBVal := util.GetDB()

	sqlVal := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", DBVal.DBHost, DBVal.DBPort, DBVal.DBUser, DBVal.DBPassword, DBVal.DBName)

	db, err = sql.Open("postgres", sqlVal)
	if err != nil {
		log.Println(err.Error())
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("DB Connected !")

}

func DBConnectTest() {
	var err error
	// DBVal := util.GetDB()
	var DBVal structures.DBConnect
	DBVal.DBHost = "localhost"
	DBVal.DBPort = "5432"
	DBVal.DBName = "postgres"
	DBVal.DBUser = "postgres"
	DBVal.DBPassword = "admin"

	sqlVal := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", DBVal.DBHost, DBVal.DBPort, DBVal.DBUser, DBVal.DBPassword, DBVal.DBName)

	db, err = sql.Open("postgres", sqlVal)
	if err != nil {
		log.Println(err.Error())
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("DB Connected !")

}
