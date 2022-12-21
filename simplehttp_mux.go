package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID   int
	Name string `json:"name"` //to lowercase
	Age  int
}

var users = []User{
	{ID: 1, Name: "GunP", Age: 23},
	{ID: 2, Name: "GunP2", Age: 24},
}

func userHandle(w http.ResponseWriter, r *http.Request) {
	u, p, ok := r.BasicAuth()
	log.Println("auth:", u, p, ok)

	if r.Method == "GET" {
		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error 500"))
			return
		}
		w.Write(b)
		return
	}
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error 500"))
			return
		}

		var u User
		err = json.Unmarshal(body, &u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error 400"))
		}

		users = append(users, u)
		fmt.Fprintf(w, "hello %s created", "POST")

		// data := []byte(`{"name":"GP001", "Method":"POST"}`) //Convert type to byte
		// w.Write(data)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// func logMIddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		start := time.Now()
// 		next.ServeHTTP(w, req)
// 		log.Printf("Server http miiddleware: %s %s %s %s", req.RemoteAddr, req.Method, req.URL, time.Since(start))
// 	}
// }

type logger struct {
	Handler http.Handler
}

func (l logger) ServeHTTP(w http.ResponseWriter, req *http.Request) { //multiplxer wrap
	start := time.Now()
	l.Handler.ServeHTTP(w, req)
	log.Printf("Server http miiddleware: %s %s %s %s", req.RemoteAddr, req.Method, req.URL, time.Since(start))

}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(401)
			w.Write([]byte(`can't parse the basic auth`))
			return
		}

		if u != "apidesign" || p != "45678" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("username/password incorrected."))
			return
		}

		fmt.Println("Authorized")
		next(w, r)
	}
}

func main() {
	

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	// mux.HandleFunc("/users", logMIddleware(userHandle)) //higher order function
	// mux.HandleFunc("/users", userHandle) //higher order function
	mux.HandleFunc("/users", AuthMiddleware(userHandle)) //

	logMux := logger{Handler: mux} //multiplxer wrap
	srv := http.Server{
		Addr:    ":2500",
		Handler: logMux, //multiplxer wrap
	}
	log.Println("Server start")
	log.Fatal(srv.ListenAndServe())
	log.Println("Exit")
}
