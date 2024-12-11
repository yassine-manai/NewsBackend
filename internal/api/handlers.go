package api

import (
	"fmt"
	"net/http"
	"news-app/internal/db"
	"news-app/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateNewsHandler handles adding a new article
func CreateNewsHandler(c *gin.Context) {
	var news db.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateNews(&news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, news)
}

// FetchNewsHandler retrieves all news
func FetchNewsHandler(c *gin.Context) {
	newsList, err := services.GetAllNews()
	//	fmt.Print(newsList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newsList)
}

// FetchNewsByIDHandler retrieves a single news article by ID
func FetchNewsByIDHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	news, err := services.GetNewsByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}

func FetchNewsByCategory(c *gin.Context) {
	cat := c.Query("category")

	if cat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category query parameter is required"})
		return
	}

	news, err := services.GetNewsByCat(cat)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}

func FetchCategories(c *gin.Context) {
	categories, err := services.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// UpdateNewsHandler updates a news article
func UpdateNewsHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	fmt.Println(id)
	fmt.Println("/////////////////")

	var updatedNews db.News
	if err := c.ShouldBindJSON(&updatedNews); err != nil {
		fmt.Println(updatedNews)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(updatedNews)

	err = services.UpdateNews(id, &updatedNews)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "news updated successfully"})
}

// DeleteNewsHandler deletes a news article
func DeleteNewsHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	err = services.DeleteNews(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "news deleted successfully"})
}
