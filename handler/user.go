package handler

import (
	"go-crowdfund/helper"
	"go-crowdfund/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct for user handler
type userHandler struct {
	userService user.Service
}

// Function to create new user handler
func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// Function to handle register user
func (h *userHandler) RegisterUser(c *gin.Context) {
	// 1. tangkap input dari user
	// 2. map input dari user ke struct RegisterUserInput
	// 3. struct di atas kita panggil di service
	// 4. service akan memanggil repository
	// 5. repository untuk simpan data user ke db

	var input user.RegisterUserInput

	// cek error dari input user
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors":errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	
	// cek error dari repository
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	
	// memanggil helper formatter
	formatter := user.FormatUser(newUser, "JWT_TOKEN")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)
	
	c.JSON(http.StatusOK, response)
}

// Function to handle login
func (h *userHandler) Login(c *gin.Context) {
	// 1. user memasukkan input (email dan password)
	// 2. input ditangkap handler
	// 3. mapping dari input user ke input struct
	// 4. input struct passing ke service
	// 5. di service mencari dengan bantuan repository user dengan email x
	// 6. mencocokan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, "JWT_TOKEN")
	response := helper.APIResponse("Successfully logged in", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}