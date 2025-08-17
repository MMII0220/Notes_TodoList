func updateCategory(c *gin.Context) {
	catID := c.Param("id")
	var input Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	var cat Category
	if err := db.First(&cat, "id = ?", catID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	cat.Name = input.Name
	db.Save(&cat)
	c.JSON(http.StatusOK, cat)
}

func deleteCategory(c *gin.Context) {
	catID := c.Param("id")
	if err := db.Delete(&Category{}, "id = ?", catID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func updateTask(c *gin.Context) {
	taskID := c.Param("id")
	var input Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	var task Task
	if err := db.First(&task, "id = ?", taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	task.Title = input.Title
	task.Description = input.Description
	task.Status = input.Status
	db.Save(&task)
	c.JSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	taskID := c.Param("id")
	if err := db.Delete(&Task{}, "id = ?", taskID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
func createCategory(c *gin.Context) {
	var input Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	input.ID = time.Now().Format("20060102150405")
	db.Create(&input)
	c.JSON(http.StatusCreated, input)
}

func getCategory(c *gin.Context) {
	catID := c.Param("id")
	var cat Category
	if err := db.First(&cat, "id = ?", catID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, cat)
}
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func createTask(c *gin.Context) {
	var input Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	input.ID = time.Now().Format("20060102150405")
	input.CreatedAt = time.Now()
	if input.Status == "" {
		input.Status = "pending"
	}
	db.Create(&input)
	c.JSON(http.StatusCreated, input)
}

func getTask(c *gin.Context) {
	taskID := c.Param("id")
	var task Task
	if err := db.First(&task, "id = ?", taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}
