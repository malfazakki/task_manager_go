package repository

import (
	"task-manager/internal/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTaskByUserID(userID uint) ([]models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(taskID uint, userID uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) CreateTask(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetTaskByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error

	return tasks, err
}

func (r *taskRepository) UpdateTask(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) DeleteTask(taskID uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{}).Error
}
