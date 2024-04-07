package routes

import (
	"net/http"
	"task-management-system/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() {
    router := mux.NewRouter()

    router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
    router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
    router.HandleFunc("/tasks/{id}", controllers.GetTaskByID).Methods("GET")
    router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}/complete", controllers.MarkTaskAsComplete).Methods("PUT")

    http.Handle("/", router)
}
