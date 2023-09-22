package handlers

import (
	"net/http"

	app "github.com/go-htmx-sample/application"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func NewDefaultHomeHandler() HomeHandler {
	return HomeHandler{}
}

func (h HomeHandler) ShowHome(c echo.Context) error {
	print("test")
	return c.Render(http.StatusOK, "index.html", nil)
}

type TaskHandler struct {
	taskService app.TaskService
}

func NewTaskHandler(taskService app.TaskService) *TaskHandler {
	return &TaskHandler{taskService}
}

// Define HTTP handlers for tasks, e.g., CreateTask, GetTask, GetAllTasks, UpdateTask, DeleteTask
// Example:
func (h *TaskHandler) CreateTask(c echo.Context) error {
	// Parse request, call service method, and handle response
	return c.JSON(http.StatusCreated, struct{}{})
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	tasks, _ := h.taskService.GetAllTasks()

	return c.JSON(http.StatusOK, tasks)
}
