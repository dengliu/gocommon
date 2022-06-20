package main

import (
	"embed"
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed public
var staticfiles embed.FS

func main() {
	var port string
	flag.StringVar(&port, "port", ":8000", "HTTP Server Port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[LOG] ", log.Lshortfile|log.LstdFlags)
	errorLog := log.New(os.Stderr, "[ERR] ", log.Lshortfile|log.LstdFlags)

	app := NewApplicatoin(infoLog, errorLog, "public")

	go app.Run()

	srv := http.Server{
		Addr:         port,
		IdleTimeout:  2 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      app.routes(),
	}

	infoLog.Printf("Listening on: http://localhost%s", port)
	if err := srv.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}
