package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
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
	posts = append(posts, post)
}
func getPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, p := range posts {
		if p.ID == id {
			post, err := json.Marshal(p)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			w.Write(post)
		}
	}
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

func main() {
	posts = []Post{{ID: "1", Title: "Hello world!", Desc: "Буквально первый пост в блоге, жесть!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!фвыфвфы"}, {ID: "2", Title: "Тест API", Desc: "Ну это типо второй пост"}}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /posts/", getPosts)
	mux.HandleFunc("POST /post/", createPost)
	mux.HandleFunc("GET /post", getPost)

	handler := corsMiddleware(mux)

	http.ListenAndServe(":8080", handler)
}
