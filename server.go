package main

import (
	"cloudbees/adapters/database"
	"cloudbees/endpoints"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func ENVSetup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error in env file" + err.Error())
	}
}

func main() {
	ENVSetup()
	port := os.Getenv("Port")
	net, err := net.Listen("tcp", ": 8080")
	if err != nil {
		log.Println(err.Error())
	}
	database.DBConnect()
	log.Println("Server started at : ", port)
	s := grpc.NewServer()
	endpoints.Endpoints(s)
	err2 := s.Serve(net)
	if err2 == nil {
		log.Println(err2.Error())
	}
}
