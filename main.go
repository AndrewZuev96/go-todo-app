package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go_proj/internal/models"
	"go_proj/internal/storage"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var store *storage.Service

func task_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

	if r.Method == http.MethodOptions {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		tasks, err := store.GetAll()
		if err != nil {
			http.Error(w, "Failed to get tasks, database query error", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tasks)

	} else if r.Method == http.MethodPost {
		var newTask models.Task
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		createTask, err := store.Create(newTask)
		if err != nil {
			http.Error(w, "Failed to create task", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createTask)

	} else if r.Method == http.MethodDelete {

		id_str := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(id_str)

		err := store.Delete(id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Task not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to delete", http.StatusInternalServerError)
			}
		}

		w.WriteHeader(http.StatusNoContent)

	} else if r.Method == http.MethodPut {

		var updated_task models.Task

		if err := json.NewDecoder(r.Body).Decode(&updated_task); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		updatedTask, err := store.Update(updated_task)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Task not found", http.StatusNotFound)
			} else {
				http.Error(w, "failed to update", http.StatusInternalServerError)
			}
			return
		}

		json.NewEncoder(w).Encode(updatedTask)

	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file found")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	server_port := os.Getenv("SERVER_PORT")
	if server_port == "" {
		server_port = ":8080"
	}

	
	conn_str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//var err error
	//connStr := "host=localhost port=5433 user=chsh password=sa12345 dbname=todo_db sslmode=disable" //todo_db=#
	//connStr := "host=127.0.0.1 port=5432 user=chsh password=sa12345 dbname=todo_db sslmode=disable"
	//connStr := "postgres://chsh:sa12345@localhost:5432/todo_db?sslmode=disable"
	db,err:= sql.Open("postgres", conn_str)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	store=storage.New(db)
	fmt.Println("Successfully connected to the databse!")
	fmt.Println("Server is running on port 8000...")
	http.HandleFunc("/tasks", task_handler)

	log.Fatal(http.ListenAndServe(server_port, nil))
}
