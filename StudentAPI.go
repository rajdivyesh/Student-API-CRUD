package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	students      []Student
	studentsMutex sync.Mutex
)

func main() {
	students = []Student{
		{ID: 1, FirstName: "Mary", LastName: "Jane"},
		{ID: 2, FirstName: "Dev", LastName: "Raj"},
	}
	http.HandleFunc("/students", listStudents)
	http.HandleFunc("/students/add", addStudent)
	http.HandleFunc("/students/update", updateStudent)
	http.HandleFunc("/students/delete", deleteStudent)
	fmt.Println("Server is running on port 8087")
	log.Fatal(http.ListenAndServe(":8087", nil))
}

func listStudents(w http.ResponseWriter, r *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	response, err := json.Marshal(students)
	if err != nil {
		http.Error(w, "Failed to Marshal JSON response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	var newStudent Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newStudent); err != nil {
		http.Error(w, "Failed to decode request body!", http.StatusBadRequest)
		return
	}
	newStudent.ID = len(students) + 1
	students = append(students, newStudent)
	response, err := json.Marshal(newStudent)
	if err != nil {
		http.Error(w, "Failed to Marshal JSON response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	var updatedStudent Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedStudent); err != nil {
		http.Error(w, "Failed to decode request body!", http.StatusBadRequest)
		return
	}
	// Find and update the student by ID
	for i, student := range students {
		if student.ID == updatedStudent.ID {
			students[i] = updatedStudent
			response, err := json.Marshal(updatedStudent)
			if err != nil {
				http.Error(w, "Failed to Marshal JSON response", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	// Find and delete the student by ID
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}
