package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
)

// func hello(w http.ResponseWriter, r *http.Request) {
// 	//	fmt.Fprintf(w, `w の実体の型: %T
// 	//
// 	// メソッド: %s
// 	// パス: %s
// 	// クエリ: %s
// 	// アクセス元: %s
// 	// ヘッダ: %v
// 	// `,
// 	//
// 	//		w,
// 	//		r.Method,
// 	//		r.URL.Path,
// 	//		r.URL.RawQuery,
// 	//		r.RemoteAddr,
// 	//		r.Header,
// 	//	)
// 	fmt.Fprintf(w, "%+v\n", r)
// }

var todoList []string

func handleTodo(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/todo.html")
	t.Execute(w, todoList)
	// fmt.Fprint(w, "aaaa")
}

func main() {
	todoList = append(todoList, "顔を洗う", "朝食を食べる", "歯を磨く")

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/todo", handleTodo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}
