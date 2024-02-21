# cloudbees

About project : 
    1. Blog backend service using grpc CRUD APIs.
    2. Implemented hexagonal architecture
    3. Used Postgresql database
    4. Error handling
    

Command to generate proto file :

protoc --go_out=invoicer --go_opt=paths=source_relative --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative invoicer.proto