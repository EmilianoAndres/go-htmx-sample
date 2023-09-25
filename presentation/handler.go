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
	return c.Render(http.StatusOK, "index.html", nil)
}

type TaskHandler struct {
	taskService app.TaskService
}

func NewTaskHandler(taskService app.TaskService) *TaskHandler {
	return &TaskHandler{taskService}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	taskName := c.FormValue("taskName")

	task, _ := h.taskService.CreateTask(taskName)

	// add custom header to trigger DOM event
	c.Response().Header().Set("HX-Trigger", "newTask")

	return c.Render(http.StatusCreated, "element.html", task)
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	tasks, _ := h.taskService.GetAllTasks()

	return c.Render(http.StatusOK, "list.html", tasks)
}
