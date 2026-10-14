package repository

import (
	"errors"
	"sync"
	"time"

	"go-backend-demo/internal/model"
)

var ErrNotFound = errors.New("task not found")

// TaskRepository: interface này tách rời logic nghiệp vụ khỏi nơi lưu trữ dữ liệu.
// Muốn đổi sang Postgres/Mongo chỉ cần viết struct khác implement interface này,
// không phải sửa service hay handler.
type TaskRepository interface {
	Create(title string) model.Task
	GetAll() []model.Task
	GetByID(id int64) (model.Task, error)
	Update(id int64, input model.UpdateTaskInput) (model.Task, error)
	Delete(id int64) error
}

// Triển khai in-memory để demo chạy ngay không cần cài database.
type inMemoryTaskRepository struct {
	mu     sync.Mutex
	tasks  map[int64]model.Task
	nextID int64
}

func NewInMemoryTaskRepository() TaskRepository {
	return &inMemoryTaskRepository{
		tasks:  make(map[int64]model.Task),
		nextID: 1,
	}
}

func (r *inMemoryTaskRepository) Create(title string) model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	task := model.Task{
		ID:        r.nextID,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	r.tasks[task.ID] = task
	r.nextID++
	return task
}

func (r *inMemoryTaskRepository) GetAll() []model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]model.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result
}

func (r *inMemoryTaskRepository) GetByID(id int64) (model.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	if !ok {
		return model.Task{}, ErrNotFound
	}
	return task, nil
}

func (r *inMemoryTaskRepository) Update(id int64, input model.UpdateTaskInput) (model.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	if !ok {
		return model.Task{}, ErrNotFound
	}

	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Done != nil {
		task.Done = *input.Done
	}

	r.tasks[id] = task
	return task, nil
}

func (r *inMemoryTaskRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return ErrNotFound
	}
	delete(r.tasks, id)
	return nil
}
