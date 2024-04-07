package models

// Task represents a task
type Task struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}
