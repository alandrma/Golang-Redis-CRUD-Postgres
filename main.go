package main

import (
	"Golang_CRUD_Postgresql/Routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := routes.Routers()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Server dijalankan pada port 8484...")

	log.Fatal(http.ListenAndServe(":8484", r))
}
