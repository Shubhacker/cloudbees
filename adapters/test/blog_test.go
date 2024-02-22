package test

import (
	"cloudbees/adapters/database"
	"cloudbees/adapters/logic"
	cloudbees "cloudbees/invoicer"
	"context"
	"log"
	"testing"
)

// Unit test cases for blog CRUD operations

func TestCreateBlog(t *testing.T) {
	log.Println("Insert API Test")
	database.DBConnectTest()
	var in cloudbees.CreatePostRequest
	in.Author = "unitTest1"
	in.Content = "unitTest1"
	in.Tags = append(in.Tags, "unitTest")
	in.Title = "UnitTest1"
	res, _ := logic.CreatePost(context.Background(), &in)

	if res.Author != in.Author {
		t.Errorf("got %q, wanted %q", res.Author, in.Author)
	}

	if res.Content != in.Content {
		t.Errorf("got %q, wanted %q", res.Content, in.Content)
	}

	if res.Title != in.Title {
		t.Errorf("got %q, wanted %q", res.Title, in.Title)
	}
}

func TestFetchBlog(t *testing.T) {
	log.Println("Read API Test")
	database.DBConnectTest()
	var in cloudbees.ReadPostRequest
	in.PostId = "c1e09269-5be6-4913-b818-e248effaf1f2"
	res, _ := logic.ReadPost(context.Background(), &in)

	if res.Author != "Author Name" {
		t.Errorf("got %q, wanted %q", res.Author, "Author Name")
	}

	if res.Content != "Contect of blog" {
		t.Errorf("got %q, wanted %q", res.Content, "Contect of blog")
	}

	if res.Title != "Title of blog" {
		t.Errorf("got %q, wanted %q", res.Title, "Title of blog")
	}
}
