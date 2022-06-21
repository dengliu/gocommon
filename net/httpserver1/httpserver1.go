package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

const addr = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/v1/hello", http.HandlerFunc(HelloHandler))
	mux.Handle("/v1/time", http.HandlerFunc(CurrentTimeHandler))
	//wrap entire mux with logger middleware
	wrappedMux := NewLogger(mux)

	log.Printf("server is listening at %s", addr)
	//use wrappedMux instead of mux as root handler
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}

//Logger is a middleware handler that does request logging
type Logger struct {
	handler http.Handler
}

//ServeHTTP handles the request by passing it to the real
//handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

//NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}


//ResponseHeader is a middleware handler that adds a header to the response
type ResponseHeader struct {
	handler http.Handler
	headerName string
	headerValue string
}

//NewResponseHeader constructs a new ResponseHeader middleware handler
func NewResponseHeader(handlerToWrap http.Handler, headerName string, headerValue string) *ResponseHeader {
	return &ResponseHeader{handlerToWrap, headerName, headerValue}
}

//ServeHTTP handles the request by adding the response header
func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//add the header
	w.Header().Add(rh.headerName, rh.headerValue)
	//call the wrapped handler
	rh.handler.ServeHTTP(w, r)
}

