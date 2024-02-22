package endpoints

import (
	"cloudbees/adapters/logic"
	cloudbees "cloudbees/invoicer"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

type MyInvoicerServer struct {
	cloudbees.UnimplementedInvoicerServer
}

func (s MyInvoicerServer) CreatePost(ctx context.Context, i *cloudbees.CreatePostRequest) (*cloudbees.PostResponse, error) {
	now := time.Now()
	// We will call all logic methods and validation
	blogData, err := logic.CreatePost(ctx, i)
	if err != nil {
		return blogData, err
	}
	log.Println("CreatePost API time took :", time.Since(now))
	return blogData, nil
}

func (s MyInvoicerServer) ReadPost(ctx context.Context, i *cloudbees.ReadPostRequest) (*cloudbees.PostResponse, error) {
	now := time.Now()
	res, err := logic.ReadPost(ctx, i)
	if err != nil {
		return &cloudbees.PostResponse{}, err
	}
	log.Println("ReadPost API time took :", time.Since(now))
	return res, nil
}

func (s MyInvoicerServer) UpdatePost(ctx context.Context, i *cloudbees.UpdatePostRequest) (*cloudbees.PostResponse, error) {
	now := time.Now()
	res, err := logic.UpdatePost(ctx, i)
	if err != nil {
		return &cloudbees.PostResponse{}, nil
	}
	log.Println("UpdatePost API time took :", time.Since(now))
	return res, nil
}

func (s MyInvoicerServer) DeletePost(ctx context.Context, i *cloudbees.DeletePostRequest) (*cloudbees.DeletePostResponse, error) {
	now := time.Now()
	res, err := logic.DeletePost(ctx, i)
	if err != nil {
		log.Println(err.Error())
		return res, err
	}
	log.Println("DeletePost API time took :", time.Since(now))
	return res, nil
}

func Endpoints(s grpc.ServiceRegistrar) {
	service := &MyInvoicerServer{}

	cloudbees.RegisterInvoicerServer(s, service)
}
