package controllers

import (
	"app/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var input models.CreateCommentReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	idUint := uint(id)

	_, err := models.GetBlogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Blog not found"})
		return
	}

	comment := models.Comment{
		AuthorName: input.Name,
		Content:    input.Content,
		PostID:     &idUint,
	}

	if err := models.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success created comment",
	})
}

func GetComments(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := models.GetBlogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Blog not found"})
		return
	}

	fmt.Println(id)

	comments, err := models.GetCommentsByPostID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   comments,
	})
}
