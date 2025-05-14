package router

import (
	"github.com/VolkHackVH/todo-list/internal/db"
	"github.com/VolkHackVH/todo-list/internal/handlers"
	"github.com/VolkHackVH/todo-list/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(db *db.Queries) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	handler := handlers.NewHandler(db)
	api := r.Group("/api")

	{
		api.POST("/register", handler.User.CreateUser)
		api.POST("/login", handler.User.LoginUser)
	}

	users := api.Group("/users")
	{
		users.GET("/:id", handler.User.GetUserByID)
		users.GET("/", handler.User.ListUsers)
		users.DELETE("/:id", handler.User.DeleteUser)
	}

	tasks := api.Group("/tasks")
	tasks.Use(middleware.AuthMiddleware())
	{
		tasks.POST("/", handler.Task.CreateTask)
		tasks.GET("/:id", handler.Task.GetTaskByID)
		tasks.DELETE("/:id", handler.Task.DeleteTask)
		tasks.PUT("/:id", handler.Task.UpdateTask)
		tasks.GET("/", handler.Task.GetAllTasks)
	}
	return r
}
