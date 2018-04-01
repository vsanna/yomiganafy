package main

import (
	"fmt"
	"net/http"
)

// ハンドラ. http.ResponseWriter, *http.Request を引数に持つ
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root %s", r.URL.Path)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

// 最もシンプルな構成
func main() {
	fmt.Println("START")
	defer fmt.Println("FINISH")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/home", homeHandler) // "/"の下でもこちらにマッチすれば処理はこちらに渡される

	http.ListenAndServe(":8080", nil)
}


func main() {
	mux := http.NewServeMux()

	// path - action(handler)を無視日つけ雨r
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/home", homeHandler)
	//mux.HandleFunc("/login", loginHandler)
	//mux.HandleFunc("/users", usersHandler)

	server := http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux
	}

	server.ListenAndServe()
}