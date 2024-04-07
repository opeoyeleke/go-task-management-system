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
    // Add sample tasks to the data store
    data.AddTask(models.Task{ID: "1", Title: "Sample Task 1", Description: "Description of Sample Task 1", Status: "todo"})
    // Add more sample tasks if needed
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

}

func TestGetTaskByID(t *testing.T) {
	setupTestEnvironment()

	req, err := http.NewRequest("GET", "/tasks/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetTaskByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(errorMessage,
			status, http.StatusOK)
	}

}

func TestUpdateTask(t *testing.T) {
	setupTestEnvironment()

	requestBody := []byte(`{"title": "Updated Task", "description": "This is an updated task", "status": "in progress"}`)
	req, err := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.UpdateTask)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(errorMessage,
			status, http.StatusOK)
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

//fix this entire file bwlow

func TestUpdateTask2(t *testing.T) {
    setupTestEnvironment()

    requestBody := []byte(`{"title": "Updated Task", "description": "This is
    an updated task", "status": "in progress"}`)
    req, err := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(requestBody))
    if err != nil {
        t.Fatal(err)
    }

    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(controllers.UpdateTask)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusNotFound {
        t.Errorf(errorMessage, status, http.StatusNotFound)
    }

}