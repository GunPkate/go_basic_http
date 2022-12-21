package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func checkHealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

type Err struct {
	Message string `json:"message"`
}

func createUsersHandler(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	users = append(users, u)

	return c.JSON(http.StatusOK, users)
}

func getUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", checkHealthHandler)

	g := e.Group("/api")
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "apidesign" && password == "45678" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/users", getUsersHandler)
	g.POST("/users", createUsersHandler)

	log.Println("Server start")
	log.Fatal(e.Start(":2500"))
	//go mod init github.com/GunPkate/echoapi load dependency
	//go mod tidy
	//go run {}
	log.Println("Exit")
}
