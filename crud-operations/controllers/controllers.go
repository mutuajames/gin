package controllers

import (
	"gin/crud-operations/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type CreateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"task"`
}

type UpdateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"task"`
}

//FindTasks GET /tasks
//Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

//CreateTask POST /tasks
//Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create task
	task := models.Task{
		AssignedTo: input.AssignedTo,
		Task:       input.Task,
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

//FindTask GET /tasks/:id
//Find a task
func FindTask(c *gin.Context) {
	// Get model if exist
	var task models.Task
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// UpdateTask PATCH /tasks/:id
//Update a task
func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&task).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

//DeleteTask DELETE /tasks/:id
//Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
