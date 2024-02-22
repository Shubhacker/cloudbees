package util

import (
	"cloudbees/adapters/structures"
	"log"
	"os"
)

// Utility or helper functions

func GetDB() structures.DBConnect {
	var DBVal structures.DBConnect
	DBVal.DBHost = os.Getenv("DBHost")
	DBVal.DBName = os.Getenv("DBName")
	DBVal.DBPassword = os.Getenv("DBPassword")
	DBVal.DBPort = os.Getenv("DBPort")
	DBVal.DBUser = os.Getenv("DBUser")
	return DBVal
}

func Error(functionName string, errMsg error) bool {
	// We can add loggers, insight logic here
	// For now using log for printing if any error occurred
	if errMsg != nil {
		log.Println("Error in : ", functionName, " Error Message : ", errMsg.Error())
		return true
	}
	return false
}
