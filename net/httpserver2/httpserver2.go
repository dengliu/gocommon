package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"time"
)

const addr = ":8080"

type JSONResponse struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}

type employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path == "/foo" {
			return
		}

		next.ServeHTTP(w, r)
		log.Println("Executing middlewareTwo again")
	})
}

func middlewareTimeout(timeout time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer func() {
			cancel()
			if ctx.Err() == context.DeadlineExceeded {
				w.WriteHeader(http.StatusGatewayTimeout)
			}
		}()

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func enforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// curl -v -X POST -H "content-type: application/json" http://localhost:8080/employee -d '{"Name":"John", "Age": 21}'
func createEmployee(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		sendResponse(w, "", fmt.Errorf("Content Type is not application/json"), http.StatusUnsupportedMediaType)
		return
	}
	var e employee
	defer r.Body.Close()
	data, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(data, &e); err != nil {
		sendResponse(w, "", err, http.StatusBadRequest)
		return
	}

	sendResponse(w, e, fmt.Errorf("fake error"), http.StatusOK)

	return
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middlewareTimeout(time.Second*5, middlewareOne(middlewareTwo(finalHandler))))
	mux.Handle("/employee", http.HandlerFunc(createEmployee))

	log.Printf("server is listening at %s", addr)
	//use wrappedMux instead of mux as root handler
	log.Fatal(http.ListenAndServe(addr, mux))
}

func sendResponse(w http.ResponseWriter, data any, err error, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	resp := JSONResponse{Data: data}

	if err != nil {
		resp.Error = err.Error()
	}

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
