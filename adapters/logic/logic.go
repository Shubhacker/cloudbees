package logic

import (
	"cloudbees/adapters/database"
	"cloudbees/adapters/util"
	cloudbees "cloudbees/invoicer"
	"context"
	"errors"
)

// All business logic we will add here

func CreatePost(ctx context.Context, i *cloudbees.CreatePostRequest) (*cloudbees.PostResponse, error) {
	blogData, err := database.AddBlog(i)
	if util.Error("AddBlog", err) {
		return blogData, err
	}
	return blogData, nil
}

func ReadPost(ctx context.Context, i *cloudbees.ReadPostRequest) (*cloudbees.PostResponse, error) {
	if i.PostId == "" {
		return &cloudbees.PostResponse{}, errors.New("Provide PostId")
	}
	res, err := database.ReadPost(i)
	if util.Error("ReadPost", err) {
		return &cloudbees.PostResponse{}, err
	}
	return res, nil
}

func UpdatePost(ctx context.Context, i *cloudbees.UpdatePostRequest) (*cloudbees.PostResponse, error) {
	if i.PostId == "" {
		return &cloudbees.PostResponse{}, errors.New("Provide PostId")
	}
	res, err := database.UpdatePost(i)
	if util.Error("UpdatePost", err) {
		return &cloudbees.PostResponse{}, nil
	}
	return res, nil
}

func DeletePost(ctx context.Context, i *cloudbees.DeletePostRequest) (*cloudbees.DeletePostResponse, error) {
	if i.PostId == "" {
		return &cloudbees.DeletePostResponse{}, errors.New("Provide PostId")
	}
	res, err := database.DeletePost(i)
	if util.Error("DeletePost", err) {
		return res, err
	}
	return res, nil
}
