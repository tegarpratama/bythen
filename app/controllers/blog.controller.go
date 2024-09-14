package controllers

import (
	"app/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	var input models.CreateBlogReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	currentUser, _ := c.Get("currentUser")
	userLoggedIn, _ := currentUser.(*models.UserLoggedIn)
	userId := uint(userLoggedIn.ID)
	userIdUint := &userId

	blog := models.Blog{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: userIdUint,
	}

	if err := models.CreateBlog(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success created blog",
	})
}

func GetBlogs(c *gin.Context) {
	search := c.Query("search")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	blogs, count, err := models.GetAllBlogs(offset, intLimit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	totalPage := int(math.Ceil(float64(count) / float64(intLimit)))

	c.JSON(http.StatusOK, gin.H{
		"status":       "ok",
		"current_page": intPage,
		"total_page":   totalPage,
		"total_data":   count,
		"data":         blogs,
	})
}

func DetailBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
		return
	}

	blog, err := models.GetBlogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   blog,
	})
}

func UpdateBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
		return
	}

	var input models.Blog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	blog, err := models.GetBlogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "blog not found"})
		return
	}

	if input.Title != "" {
		blog.Title = input.Title
	}

	if input.Content != "" {
		blog.Content = input.Content
	}

	if err := models.UpdateBlog(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   blog,
	})
}

func DeleteBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
		return
	}

	_, err = models.GetBlogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Product not found"})
		return
	}

	if err := models.DeleteBlog(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
