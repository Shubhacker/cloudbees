package controller

import (
	cloudbees "cloudbees/invoicer"
	"context"
)

type EndpointInterface interface {
	CreatePost(ctx context.Context, i *cloudbees.CreatePostRequest) (*cloudbees.PostResponse, error)
}
