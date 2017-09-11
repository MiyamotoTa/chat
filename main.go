package main

import (
	"flag"
	"github.com/MiyamotoTa/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
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
	var googleClientId = flag.String("google_client_id", "", "Google OAuth client ID")
	var googleClientSecret = flag.String("google_client_secret", "", "Google OAuth client secret")
	flag.Parse()
	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey("miyamoto_ta/chat")
	gomniauth.WithProviders(
		google.New(
			*googleClientId,
			*googleClientSecret,
			"http://localhost:8080/auth/callback/google",
		),
	)
	r := newRoom()
	// 記録を無効化
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	// チャットルームを開始
	go r.run()
	// WEBサーバーを起動する
	log.Println("Webサーバ起動します ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Listen And Serve:", err)
	}
}
