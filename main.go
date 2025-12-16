package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func task_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

	if r.Method == http.MethodOptions {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		rows, err := db.Query("SELECT id, title, completed FROM tasks ORDER BY id")
		if err != nil {
			http.Error(w, "database query error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var task []Task

		for rows.Next() {
			var t Task
			if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
				continue
			}
			task = append(task, t)
		}

		json.NewEncoder(w).Encode(task)

	} else if r.Method == http.MethodPost {
		var newTask Task
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		query := "INSERT INTO tasks (title,completed) VALUES ($1,$2) RETURNING id"

		err := db.QueryRow(query, newTask.Title, newTask.Completed).Scan(&newTask.ID)
		if err != nil {
			http.Error(w, "database save error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)

	} else if r.Method == http.MethodDelete {

		id_str := r.URL.Query().Get("id")
		if id_str == "" {
			http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
			return
		}

		result, err := db.Exec("DELETE FROM tasks WHERE id = $1", id_str)
		if err != nil {
			http.Error(w, "Database delete error", http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	} else if r.Method == http.MethodPut {

		var updated_task Task

		if err := json.NewDecoder(r.Body).Decode(&updated_task); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if updated_task.ID == 0 {
			http.Error(w, "task ID is required", http.StatusBadRequest)
			return
		}

		res, err := db.Exec("UPDATE tasks SET title=$1, completed=$2 WHERE id=$3",
			updated_task.Title, updated_task.Completed, updated_task.ID)

		if err != nil {
			http.Error(w, "Database update error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			http.Error(w, "Task not fount", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(updated_task)

	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	fmt.Println("Hello world")
	var err error
	//connStr := "host=localhost port=5433 user=chsh password=sa12345 dbname=todo_db sslmode=disable" //todo_db=#
	connStr := "host=127.0.0.1 port=5432 user=chsh password=sa12345 dbname=todo_db sslmode=disable"
	//connStr := "postgres://chsh:sa12345@localhost:5432/todo_db?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	fmt.Println("Successfully connected to the databse!")
	fmt.Println("Server is running on port 8000...")
	http.HandleFunc("/tasks", task_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
