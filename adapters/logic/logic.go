package logic

import (
	"cloudbees/adapters/database"
	cloudbees "cloudbees/invoicer"
	"context"
	"errors"
)

func CreatePost(ctx context.Context, i *cloudbees.CreatePostRequest) (*cloudbees.PostResponse, error) {
	// All business logic we can add here
	blogData, err := database.AddBlog(i)
	if err != nil {
		return blogData, err
	}
	return blogData, nil
}

func ReadPost(ctx context.Context, i *cloudbees.ReadPostRequest) (*cloudbees.PostResponse, error) {
	if i.PostId == "" {
		return &cloudbees.PostResponse{}, errors.New("Provide PostId")
	}
	res, err := database.ReadPost(i)
	if err != nil {
		return &cloudbees.PostResponse{}, err
	}
	return res, nil
}

func UpdatePost(ctx context.Context, i *cloudbees.UpdatePostRequest) (*cloudbees.PostResponse, error) {
	if i.PostId == "" {
		return &cloudbees.PostResponse{}, errors.New("Provide PostId")
	}
	res, err := database.UpdatePost(i)
	if err != nil {
		return &cloudbees.PostResponse{}, nil
	}
	return res, nil
}

func DeletePost(ctx context.Context, i *cloudbees.DeletePostRequest) (*cloudbees.DeletePostResponse, error) {
	if i.PostId == "" {
		return &cloudbees.DeletePostResponse{}, errors.New("Provide PostId")
	}
	res, err := database.DeletePost(i)
	if err != nil {
		return res, err
	}
	return res, nil
}
