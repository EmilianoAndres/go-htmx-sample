package router

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-htmx-sample/application"
	"github.com/go-htmx-sample/infrastructure"
	handlers "github.com/go-htmx-sample/presentation"
	"github.com/labstack/echo/v4"
)

// CreateRouter creates a GIN router
func CreateRouter() *echo.Echo {
	router := echo.New()

	// extra configuration

	configureMappings(router)

	return router
}

func configureMappings(e *echo.Echo) {

	// Template Parsing by domain
	t := &Template{
		templates: parseTemplates(),
	}

	// set Renderer with template parsing
	e.Renderer = t

	// Repository
	taskRepository := infrastructure.NewDefaultTaskRepository()
	// Services
	taskService := application.NewTaskService(taskRepository)
	// Handlers
	taskHandler := handlers.NewTaskHandler(taskService)
	homeHandler := handlers.NewDefaultHomeHandler()

	e.GET("/", homeHandler.ShowHome)

	// tasks
	e.GET("/tasks", taskHandler.GetAllTasks)
	e.POST("/create", taskHandler.CreateTask)
}

func parseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("./public/views", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}
