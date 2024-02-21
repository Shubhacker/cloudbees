package repository

import (
	cloudbees "cloudbees/invoicer"
	"context"
	"log"
)

type CloudbeesInterface interface {
	CreatePost(ctx context.Context, i *cloudbees.CreatePostRequest) (*cloudbees.PostResponse, error)
}

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (A Adapter) CreatePost(ctx context.Context, i *cloudbees.CreatePostRequest) (*cloudbees.PostResponse, error) {
	log.Println("In repo")
	return &cloudbees.PostResponse{}, nil
}

func TestFun(c CloudbeesInterface) {
	// c.CreatePost()

}
