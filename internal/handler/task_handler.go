package handler

import (
	"net/http"
	"task-manager/internal/models"
	"task-manager/internal/usecase"

	"github.com/labstack/echo"
)

type TaskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var task models.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	userID := c.Get("userID").(uint)
	task.UserID = userID

	if err := h.taskUsecase.CreateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	userID := c.Get("userID").(uint)

	tasks, err := h.taskUsecase.GetTasks(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, tasks)
}
