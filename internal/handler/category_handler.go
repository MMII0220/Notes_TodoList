package handler

import (
	"Notes_TodoList/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	usecase domain.CategoryUsecase
}

func NewCategoryHandler(router *gin.Engine, uc *domain.CategoryUsecase) {
	handler := &CategoryHandler{usecase: *uc}

	group := router.Group("/categories")
	{
		group.GET("/ping", pingHandler)

		group.POST("/", handler.Create)
		group.GET("/:id", handler.Get)
		group.GET("/", handler.GetAll)
		group.PUT("/:id", handler.Update)
		group.DELETE("/:id", handler.Delete)
	}
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func (handler *CategoryHandler) Create(c *gin.Context) {
	var myCategory domain.Category

	if err := c.ShouldBindJSON(&myCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	created, err := handler.usecase.Create(myCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании"})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (handler *CategoryHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	myCategory, err := handler.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найдена"})
		return
	}

	c.JSON(http.StatusOK, myCategory)
}

func (handler *CategoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	var myCategory domain.Category
	if err := c.ShouldBindJSON(&myCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	updated, err := handler.usecase.Update(id, myCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (handler *CategoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	err = handler.usecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Категория удалена"})
}

func (handler *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := handler.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}
