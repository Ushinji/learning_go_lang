package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // wに入力した内容がクライアントへ出力される
}

func main() {
	http.HandleFunc("/", sayhelloName)				// アクセスルーティングの定義
	err := http.ListenAndServe(":9090", nil) 	// 監視するポートの設定
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
