package util

import (
	"cloudbees/adapters/structures"
	"os"
)

func GetDB() structures.DBConnect {
	var DBVal structures.DBConnect
	DBVal.DBHost = os.Getenv("DBHost")
	DBVal.DBName = os.Getenv("DBName")
	DBVal.DBPassword = os.Getenv("DBPassword")
	DBVal.DBPort = os.Getenv("DBPort")
	DBVal.DBUser = os.Getenv("DBUser")
	return DBVal
}
