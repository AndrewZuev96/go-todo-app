package storage

import (
	"database/sql"
	"fmt"
	"go_proj/internal/models"
	"time"
)

type Service struct {
	db *sql.DB
}

func New(db *sql.DB) *Service {
	return &Service{db: db}
}

// ////////////////////////
// 1. Получить все задачи
// ///////////////////////
func (s *Service) GetAll() ([]models.Task, error) {
	rows, err := s.db.Query("SELECT id,title,completed FROM tasks ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
			continue
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// ////////////////////////
// 2. Создать задачу
// ///////////////////////
func (s *Service) Create(task models.Task) (models.Task, error) {
	query := "INSERT INTO tasks (title,completed) VALUES ($1,$2) RETURNING id"
	err := s.db.QueryRow(query, task.Title, task.Completed).Scan(&task.ID)
	return task, err
}

// ////////////////////////
// 3. Обновить задачу
// ///////////////////////
func (s *Service) Update(task models.Task) (models.Task, error) {
	res, err := s.db.Exec("UPDATE tasks SET title=$1, completed=$2 WHERE id=$3",
		task.Title, task.Completed, task.ID)

	if err != nil {
		return task, err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return task, sql.ErrNoRows
	}
	return task, nil
}

// ////////////////////////
// 4. Удалить задачу
// ///////////////////////
func (s *Service) Delete(id int) error {
	res, err := s.db.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func SendNotification(taskTitle string) {
	fmt.Printf("start sending email for task: '%s'...\n", taskTitle)
	time.Sleep(5 * time.Second)
	fmt.Printf("Email sent for task: '%s'!\n", taskTitle)
}
