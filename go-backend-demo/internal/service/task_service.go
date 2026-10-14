package service

import (
	"go-backend-demo/internal/model"
	"go-backend-demo/internal/repository"
)

// Service chứa business logic, không biết gì về HTTP (không import Gin ở đây).
type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(input model.CreateTaskInput) model.Task {
	return s.repo.Create(input.Title)
}

func (s *TaskService) ListTasks() []model.Task {
	return s.repo.GetAll()
}

func (s *TaskService) GetTask(id int64) (model.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) UpdateTask(id int64, input model.UpdateTaskInput) (model.Task, error) {
	return s.repo.Update(id, input)
}

func (s *TaskService) DeleteTask(id int64) error {
	return s.repo.Delete(id)
}
