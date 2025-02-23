package handler

import (
	"go-crowdfund/helper"
	"go-crowdfund/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita panggil di service
	// service akan memanggil repository
	// repository untuk simpan data user ke db

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