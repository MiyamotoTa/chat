package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"flag"
)

type templateHandler struct {
	once     sync.Once
	filename string
	template *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.template = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.template.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "Application address")
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルームを開始
	go r.run()
	// WEBサーバーを起動する
	log.Println("Webサーバ起動します ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Listen And Serve:", err)
	}
}
