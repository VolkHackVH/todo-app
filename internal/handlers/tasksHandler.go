package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/VolkHackVH/todo-list/internal/db"
	"github.com/VolkHackVH/todo-list/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type TasksHandler struct {
	Db *service.TaskService
}

func newTaskHandler(db *db.Queries) *TasksHandler {
	return &TasksHandler{
		Db: service.NewTaskService(db),
	}
}

func (h *TasksHandler) CreateTask(c *gin.Context) {
	var request struct {
		Text string `json:"text" binding:"required"`
	}

	time := pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.Db.CreateTask(c.Request.Context(), request.Text, time)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TasksHandler) GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task id not found"})
		return
	}

	task, err := h.Db.GetTaskByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TasksHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task id not found"})
		return
	}

	if err := h.Db.DeleteTask(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"success": "task deleted"})
}

func (h *TasksHandler) UpdateTask(c *gin.Context) {
	var request struct {
		Text string `json:"text" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task id not found"})
		return
	}

	if err := h.Db.UpdateTask(c.Request.Context(), id, request.Text); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "update complited"})
}

func (h *TasksHandler) GetAllTasks(c *gin.Context) {
	task, err := h.Db.GetAllTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch"})
		return
	}

	c.JSON(http.StatusOK, task)
}
