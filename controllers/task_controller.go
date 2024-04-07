package controllers

import (
	"encoding/json"
	"net/http"

	"task-management-system/data"
	"task-management-system/models"
	"task-management-system/utils"

	"github.com/google/uuid"
)

// taskNotFoundErrorMessage is the error message returned when a task is not found
const taskNotFoundErrorMessage = "Task not found"

// CreateTask creates a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    task.ID = uuid.New().String()

    // Validate the task
    if err := utils.ValidateTask(task); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, err.Error())
        return
    }

    // Add the task to the data store
    data.AddTask(task)
    utils.RespondWithJSON(w, http.StatusCreated, task)
}

// GetTasks returns all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
    tasks := data.GetTasks()
    utils.RespondWithJSON(w, http.StatusOK, tasks)
}


// GetTaskByID returns a task by ID
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
    id := utils.GetIDFromRequest(r)
    task, err := data.GetTaskByID(id)
    if err != nil {
        utils.RespondWithError(w, http.StatusNotFound, taskNotFoundErrorMessage)
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, task)
}

// UpdateTask updates a task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
    id := utils.GetIDFromRequest(r)
    if id == "" {
        utils.RespondWithError(w, http.StatusBadRequest, "Task ID is required")
        return
    }

    // Decode the incoming task
    var updatedTask models.Task
    err := json.NewDecoder(r.Body).Decode(&updatedTask)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Ensure the task ID in the request payload matches the URL
    if updatedTask.ID != id {
        utils.RespondWithError(w, http.StatusBadRequest, "Task ID in request payload does not match URL")
        return
    }

    // Validate the task
    if err := utils.ValidateTask(updatedTask); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, err.Error())
        return
    }

    // Update the task
    task, err := data.UpdateTask(id, updatedTask)
    if err != nil {
        utils.RespondWithError(w, http.StatusNotFound, taskNotFoundErrorMessage)
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, task)
}

// MarkTaskAsComplete marks a task as complete
func MarkTaskAsComplete(w http.ResponseWriter, r *http.Request) {
    id := utils.GetIDFromRequest(r)
    task, err := data.MarkTaskAsComplete(id)
    if err != nil {
        utils.RespondWithError(w, http.StatusNotFound, taskNotFoundErrorMessage)
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, task)
}
