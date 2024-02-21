package endpoints

import (
	"cloudbees/adapters/database"
	cloudbees "cloudbees/invoicer"
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc"
)

type MyInvoicerServer struct {
	cloudbees.UnimplementedInvoicerServer
}

func (s MyInvoicerServer) CreatePost(ctx context.Context, i *cloudbees.CreatePostRequest) (*cloudbees.PostResponse, error) {
	now := time.Now()
	blogData, err := database.AddBlog(i)
	if err != nil {
		return blogData, err
	}
	log.Println("CreatePost API time took :", time.Since(now))
	return blogData, nil
}

func (s MyInvoicerServer) ReadPost(ctx context.Context, i *cloudbees.ReadPostRequest) (*cloudbees.PostResponse, error) {
	now := time.Now()
	if i.PostId == "" {
		return &cloudbees.PostResponse{}, errors.New("Provide PostId")
	}
	res, err := database.ReadPost(i)
	if err != nil {
		return &cloudbees.PostResponse{}, err
	}
	log.Println("ReadPost API time took :", time.Since(now))
	return res, nil
}

func (s MyInvoicerServer) UpdatePost(ctx context.Context, i *cloudbees.UpdatePostRequest) (*cloudbees.PostResponse, error) {
	now := time.Now()
	if i.PostId == "" {
		return &cloudbees.PostResponse{}, errors.New("Provide PostId")
	}
	res, err := database.UpdatePost(i)
	if err != nil {
		return &cloudbees.PostResponse{}, nil
	}
	log.Println("UpdatePost API time took :", time.Since(now))
	return res, nil
}

func Endpoints(s grpc.ServiceRegistrar) {
	service := &MyInvoicerServer{}

	cloudbees.RegisterInvoicerServer(s, service)
}
