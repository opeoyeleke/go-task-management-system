package data

import (
	"errors"
	"sync"
	"task-management-system/models"
)

const taskNotFoundErrorMessage = "task not found"


var (
    mutex sync.Mutex
    tasks []models.Task
)

// AddTask adds a new task to the data store
func AddTask(task models.Task) {
    mutex.Lock()
    defer mutex.Unlock()
    tasks = append(tasks, task)
}

// GetTasks returns all tasks
func GetTasks() []models.Task {
    return tasks
}


// GetTaskByID returns a task by ID
func GetTaskByID(id string) (models.Task, error) {
    for _, task := range tasks {
        if task.ID == id {
            return task, nil
        }
    }
    return models.Task{}, errors.New(taskNotFoundErrorMessage)
}

// UpdateTask updates a task
func UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i] = updatedTask
            return tasks[i], nil
        }
    }
    return models.Task{}, errors.New(taskNotFoundErrorMessage)
}

// MarkTaskAsComplete marks a task as complete
func MarkTaskAsComplete(id string) (models.Task, error) {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Status = "completed"
            return tasks[i], nil
        }
    }
    return models.Task{}, errors.New(taskNotFoundErrorMessage)
}
