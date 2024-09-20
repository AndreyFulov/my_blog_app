package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
}

type DataBase struct {
}

var dbInfo string

func (d *DataBase) InitInfo(host, port, user, password, dbname, sslmode string) {
	dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

}

func (d *DataBase) CreateTable() error {

	//Подключаемся к БД
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
		Id SERIAL PRIMARY KEY,
		Title TEXT,
		Description TEXT
	);`); err != nil {
		return err
	}
	return nil
}

func (d *DataBase) CreatePost(p Post) error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	data := `INSERT INTO posts (Title, Description) VALUES ($1, $2)`
	if _, err = db.Exec(data, p.Title, p.Description); err != nil {
		return err
	}
	return nil
}
func (d DataBase) GetAllPosts() ([]Post, error) {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Title, &post.Description)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (d *DataBase) GetPostByID(id int64) (Post, error) {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return Post{}, err
	}
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM posts WHERE ID = $1;`, id)
	if err != nil {
		return Post{}, err
	}
	var post Post
	for rows.Next() {
		err := rows.Scan(&post.Id, &post.Title, &post.Description)
		if err != nil {
			return Post{}, nil
		}
	}
	return post, nil
}
