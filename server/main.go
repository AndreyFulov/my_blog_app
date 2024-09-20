package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := db.GetAllPosts()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	p, err := json.Marshal(posts)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	w.Write(p)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	var post Post
	err = json.Unmarshal(b, &post)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	db.CreatePost(post)
}
func getPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	id64 := int64(id)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	posts, err := db.GetPostByID(id64)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	post, err := json.Marshal(posts)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	w.Write(post)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем запросы с любых источников
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Разрешаем следующие методы запроса
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		// Разрешаем определенные заголовки
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Если это предварительный запрос OPTIONS, то не передаем его дальше
		if r.Method == "OPTIONS" {
			return
		}

		// Передаем запрос дальше по цепочке middleware
		next.ServeHTTP(w, r)
	})
}

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")
var sslmode = os.Getenv("SSLMODE")
var db DataBase

func main() {
	db.InitInfo(host, port, user, password, dbname, sslmode)
	time.Sleep(5 * time.Second)
	fmt.Printf("Сервер запущен!")
	db.CreateTable()
	time.Sleep(5 * time.Second)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /posts/", getPosts)
	mux.HandleFunc("POST /post/", createPost)
	mux.HandleFunc("GET /post", getPost)

	handler := corsMiddleware(mux)

	http.ListenAndServe(":8080", handler)
}
