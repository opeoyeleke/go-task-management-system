package main

import (
	"log"
	"net/http"
	"task-management-system/routes"
)

func main() {
    routes.InitRoutes()

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
