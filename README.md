# cloudbees

About project : 
    1. Blog backend service using grpc CRUD APIs.
    2. Used Postgresql database
    3. Error handling
    

Command to generate proto file :

protoc --go_out=invoicer --go_opt=paths=source_relative --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative invoicer.proto

Project strucutre :

Main File -> server.go

Required tables insert query -> DBTables

API requests -> API_Requests

ENV variables -> .env

grpc Auto generated files -> invoicer

proto file for grpc -> invoicer.proto

Methods for APIs -> endpoints -> endpoints.go

Logic -> adapters -> logic.go

DBCall -> adapters -> database.go

Test cases -> adapters -> test -> blog_test.go
