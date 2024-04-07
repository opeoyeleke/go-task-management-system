package data

import (
	"errors"
	"sync"
	"task-management-system/models"
)

var (
    mutex sync.Mutex
    tasks []models.Task
)

func AddTask(task models.Task) {
    mutex.Lock()
    defer mutex.Unlock()
    tasks = append(tasks, task)
}

func GetTasks() []models.Task {
    return tasks
}

const taskNotFoundErrorMessage = "task not found"

func GetTaskByID(id string) (models.Task, error) {
    for _, task := range tasks {
        if task.ID == id {
            return task, nil
        }
    }
    return models.Task{}, errors.New(taskNotFoundErrorMessage)
}

func UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i] = updatedTask
            return tasks[i], nil
        }
    }
    return models.Task{}, errors.New(taskNotFoundErrorMessage)
}

func MarkTaskAsComplete(id string) (models.Task, error) {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Status = "completed"
            return tasks[i], nil
        }
    }
    return models.Task{}, errors.New(taskNotFoundErrorMessage)
}
