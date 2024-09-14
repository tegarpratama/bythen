package controllers

import (
	"app/config"
	"app/helpers"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input models.Register
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	customerExist := models.GetCustomerByEmail(input.Email)
	if customerExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Email already registered"})
		return
	}

	if input.Password != input.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Passwords not match"})
		return
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	user := models.User{
		Name:          input.Name,
		Email:         input.Email,
		Password_hash: hashedPassword,
	}

	if err := models.CreateCustomer(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": &models.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

func Login(c *gin.Context) {
	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	userExist := models.GetCustomerByEmail(input.Email)
	if userExist.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Wrong email or password"})
		return
	}

	userLoggedIn := &models.UserLoggedIn{
		ID:    int(userExist.ID),
		Name:  userExist.Name,
		Email: userExist.Email,
	}

	if err := helpers.VerifyPassword(userExist.Password_hash, input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid email or Password"})
		return
	}

	access_token, err := helpers.CreateToken(config.ENV.AccessTokenExpiresIn, userLoggedIn, config.ENV.AccessTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	userLoggedIn.Token = access_token

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   userLoggedIn,
	})
}
