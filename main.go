// main.go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	connStr := "host=db user=dev password=dev sslmode=disable"
	// connStr := "user=dev password=dev sslmode=disable"

	db, err := sqlx.Open("postgres", connStr)

	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("postgres connected :: %#v \n", connStr)
	}

	tmpl := template.Must(template.ParseGlob("./app/templates/*.tmpl"))
	fmt.Printf("template found :: %#v \n", tmpl.Name())

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("dir :: ", dir)

	files, err := ioutil.ReadDir("./app")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	data, err := ioutil.ReadFile("./app/resources/apptest.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	} else {
		fmt.Println("reading contents of file:", string(data))
	}
	r := mux.NewRouter()

	server := &http.Server{
		Addr: ":8080",
	}

	r.HandleFunc("/", hello)
	http.Handle("/", r)

	fmt.Printf("starting server :: %#v \n", server.Addr)

	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))

}

func hello(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "success",
		Message: "Hello world!",
	}

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)

}
