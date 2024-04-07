package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"task-management-system/models"

	"github.com/gorilla/mux"
)

// RespondWithError responds with an error message.
func RespondWithError(w http.ResponseWriter, code int, msg string) {
    RespondWithJSON(w, code, map[string]string{"error": msg})
}

// RespondWithJSON responds with JSON data.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

// ValidateTask validates a task object.
func ValidateTask(task models.Task) error {
    if task.Title == "" {
        return errors.New("Title is required")
    }
    if task.Description == "" {
        return errors.New("Description is required")
    }
    if task.Status != "todo" && task.Status != "in progress" && task.Status != "completed" {
        return errors.New("Invalid status")
    }
    return nil
}

// GetIDFromRequest extracts the ID parameter from the request URL.
func GetIDFromRequest(r *http.Request) string {
    vars := mux.Vars(r)
    id := vars["id"]
    return id
}