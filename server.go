package main

import (
	"cloudbees/adapters/database"
	"cloudbees/adapters/util"
	"cloudbees/endpoints"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func ENVSetup() {
	err := godotenv.Load(".env")
	util.Error("ENVSetup", err)
}

func main() {
	ENVSetup()
	port := os.Getenv("Port")
	net, err := net.Listen("tcp", ": 8080")
	if util.Error("net.Listen", err) {
		return
	}
	database.DBConnect()
	log.Println("Server started at : ", port)
	s := grpc.NewServer()
	endpoints.Endpoints(s)
	err2 := s.Serve(net)
	if util.Error("s.Server", err2) {
		return
	}
}
