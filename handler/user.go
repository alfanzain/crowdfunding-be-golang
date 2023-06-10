package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse(
			"Register account failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.Register(input)
	if err != nil {
		response := helper.ApiResponse(
			"Register account failed",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")

	response := helper.ApiResponse(
		"Account has been registered",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukkan input berupa email dan password
	// input ditangkap handler
	// mapping dari input user ke input struct
	// input struct passing service
	// di service mencari dengan bantuan repository user dengan email x
	// mencocokkan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse(
			"Login failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedOnUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse(
			"Login failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedOnUser, "tokentokentoken")

	response := helper.ApiResponse(
		"Login successful",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}
