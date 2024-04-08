package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-management-system/controllers"
	"task-management-system/data"
	"task-management-system/models"
	"testing"
)

const errorMessage = "handler returned wrong status code: got %v want %v"

func setupTestEnvironment() {
	// Create a new task
	task1 := models.Task{ID: "1", Title: "Sample Task 1", Description: "Description of Sample Task 1", Status: "todo"}
	data.AddTask(task1)
	
	// Create another task
	task2 := models.Task{ID: "2", Title: "Sample Task 2", Description: "Description of Sample Task 2", Status: "in progress"}
	data.AddTask(task2)	
}

func TestCreateTask(t *testing.T) {
	setupTestEnvironment()

	requestBody := []byte(`{"title": "Test Task", "description": "This is a test task", "status": "todo"}`)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controllers.CreateTask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf(errorMessage, status, http.StatusCreated)
	}

	expected := models.Task{Title: "Test Task", Description: "This is a test task", Status: "todo"}
	var actual models.Task
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}
	if actual.Title != expected.Title || actual.Description != expected.Description || actual.Status != expected.Status {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetTasks(t *testing.T) {
	setupTestEnvironment()

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetTasks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(errorMessage,
			status, http.StatusOK)
	}

	var tasks []models.Task
	if err := json.Unmarshal(rr.Body.Bytes(), &tasks); err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 2 {
		t.Errorf("handler returned unexpected number of tasks: got %v want %v",
			len(tasks), 2)
	}

}

// TestGetTaskByID tests the GetTaskByID handler function
func TestGetTaskByID(t *testing.T) {
	setupTestEnvironment()
	
	// Create a request to retrieve the task with ID 1
	req, err := http.NewRequest("GET", "/tasks/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function to handle the request
	handler := http.HandlerFunc(controllers.GetTaskByID)
	handler.ServeHTTP(rr, req)

	// Check if the response status code is OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response body contains the expected task information
	expected := `{"id":"1","title":"Sample Task 1","description":"Description of Sample Task 1","status":"todo"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestUpdateTask(t *testing.T) {
	setupTestEnvironment()

	requestBody := []byte(`{"id": "1", "title": "Updated Task", "description": "This is an updated task", "status": "in progress"}`)
	req, err := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controllers.UpdateTask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(errorMessage, status, http.StatusOK)
	}

	expected := models.Task{ID: "1", Title: "Updated Task", Description: "This is an updated task", Status: "in progress"}
	var actual models.Task
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}
	if actual.Title != expected.Title || actual.Description != expected.Description || actual.Status != expected.Status {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestMarkTaskAsComplete(t *testing.T) {
	setupTestEnvironment()

	req, err := http.NewRequest("PUT", "/tasks/1/complete", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.MarkTaskAsComplete)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(errorMessage,
			status, http.StatusOK)
	}

}

