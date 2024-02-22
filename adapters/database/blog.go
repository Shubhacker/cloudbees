package database

import (
	cloudbees "cloudbees/invoicer"
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
)

func AddBlog(i *cloudbees.CreatePostRequest) (*cloudbees.PostResponse, error) {
	var inputArgs []interface{}
	var blogData cloudbees.PostResponse
	var tags string
	sqlQuery := `INSERT INTO public.blog (title, content, author, tags, publicationDate) VALUES (
		?, ?, ?, ?, current_timestamp) returning blogId, content, author, tags, publicationDate, title;`
	inputArgs = append(inputArgs, i.Title)
	inputArgs = append(inputArgs, i.Content)
	dt, err3 := json.Marshal(i.Tags)
	if err3 != nil {
		log.Println(err3.Error())
	}
	inputArgs = append(inputArgs, i.Author)
	inputArgs = append(inputArgs, string(dt))

	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
	err := db.QueryRow(sqlQuery, inputArgs...).Scan(&blogData.PostId, &blogData.Content, &blogData.Author, &tags, &blogData.PublicationDate, &blogData.Title)
	if err != nil {
		log.Println(err)
		return &cloudbees.PostResponse{}, err
	}

	blogData.Tags = append(blogData.Tags, tags)

	return &blogData, nil
}

func ReadPost(i *cloudbees.ReadPostRequest) (*cloudbees.PostResponse, error) {
	var response cloudbees.PostResponse
	var tags string
	sqlQuery := `select blogid, title, "content", author, publicationDate, tags from public.blog where blogid = $1`

	err := db.QueryRow(sqlQuery, i.PostId).Scan(&response.PostId, &response.Title, &response.Content, &response.Author, &response.PublicationDate, &tags)
	if err != nil {
		log.Println(err.Error())
	}
	response.Tags = append(response.Tags, tags)

	return &response, nil
}

func UpdatePost(i *cloudbees.UpdatePostRequest) (*cloudbees.PostResponse, error) {
	var response cloudbees.PostResponse
	var inputArgs []interface{}
	var tags string
	var isFirst = true
	var sqlVal string
	sqlQuery := `update public.blog set `

	if i.Title != "" {
		sqlVal = `title = ?`
		if !isFirst {
			sqlVal = `, title = ?`
		}
		sqlQuery += sqlVal
		isFirst = false
		inputArgs = append(inputArgs, i.Title)
	}

	if i.Content != "" {
		sqlVal = `"content" = ?`
		if !isFirst {
			sqlVal = `, "content" = ?`
		}
		sqlQuery += sqlVal
		isFirst = false
		inputArgs = append(inputArgs, i.Content)
	}

	if i.Author != "" {
		sqlVal = `author = ?`
		if !isFirst {
			sqlVal = `, author = ?`
		}
		sqlQuery += sqlVal
		isFirst = false
		inputArgs = append(inputArgs, i.Author)
	}

	if i.Tags != nil {
		sqlVal = `tags = ?`
		if !isFirst {
			sqlVal = `, tags = ?`
		}
		sqlQuery += sqlVal
		isFirst = false
		dt, err3 := json.Marshal(i.Tags)
		if err3 != nil {
			log.Println(err3.Error())
		}
		inputArgs = append(inputArgs, dt)
	}
	sqlQuery += ` where blogid = ? returning blogId, content, author, tags, publicationDate, title;`
	inputArgs = append(inputArgs, i.PostId)

	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)

	err := db.QueryRow(sqlQuery, inputArgs...).Scan(&response.PostId, &response.Content, &response.Author, &tags, &response.PublicationDate, &response.Title)
	if err != nil {
		log.Println(err.Error())
		return &cloudbees.PostResponse{}, err
	}
	response.Tags = append(response.Tags, tags)

	return &response, nil
}

func DeletePost(i *cloudbees.DeletePostRequest) (*cloudbees.DeletePostResponse, error) {
	var response cloudbees.DeletePostResponse
	var blogId string
	sqlQuery := `delete from public.blog where blogid = $1 returning blogid`

	err := db.QueryRow(sqlQuery, i.PostId).Scan(&blogId)
	if err != nil {
		log.Println(err.Error())
		response.IsDeleted = "Failed to delete the post !"
		return &response, err
	}
	response.IsDeleted = "Post deleted successfully !"

	return &response, nil
}
