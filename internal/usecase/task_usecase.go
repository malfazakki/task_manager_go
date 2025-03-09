package usecase

import (
	"errors"
	"task-manager/internal/models"
	"task-manager/internal/repository"
)

type TaskUsecase interface {
	CreateTask(task *models.Task) error
	GetTasks(userID uint) ([]models.Task, error)
	UpdateTask(task *models.Task, userID uint) error
	DeleteTask(taskID uint, userID uint) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo}
}

func (u *taskUsecase) CreateTask(task *models.Task) error {
	return u.taskRepo.CreateTask(task)
}

func (u *taskUsecase) GetTasks(userID uint) ([]models.Task, error) {
	return u.taskRepo.GetTaskByUserID(userID)
}
func (u *taskUsecase) UpdateTask(task *models.Task, userID uint) error {
	if task.UserID != userID {
		return errors.New("unauthorized")
	}

	return u.taskRepo.UpdateTask(task)
}

func (u *taskUsecase) DeleteTask(taskID uint, userID uint) error {
	return u.taskRepo.DeleteTask(taskID, userID)
}
