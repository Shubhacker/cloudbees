package service

import (
	"cloudbees/controller"
	cloudbees "cloudbees/invoicer"
	"context"
	"log"
)

type Adapter struct {
	cb controller.EndpointInterface
}

func NewAdapter(port controller.EndpointInterface) *Adapter {
	return &Adapter{
		cb: port,
	}
}

func (api Adapter) CreatePost(ctx context.Context, i *cloudbees.CreatePostRequest) (*cloudbees.PostResponse, error) {
	res, err := api.cb.CreatePost(ctx, i)
	log.Println("In Service")
	return res, err
}
