package handler

import (
	"SalimProject/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		logrus.Printf("Invalid input: %v", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Auth.CreateUser(input)
	if err != nil {
		logrus.Printf("User registration failed: %v", err.Error())
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Printf("User with id: %v just registered", id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password_hash" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.
			StatusBadRequest, err.Error())
		return
	}
	tokenA, tokenR, err := h.services.Auth.GenerateToken(input.Username, input.Password, input.Email)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  tokenA,
		"refresh_token": tokenR,
	})
}

type refreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (h *Handler) refreshHandler(c *gin.Context) {
	var input refreshInput
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.
			StatusBadRequest, err.Error())
		return
	}
	userId, err := h.services.Auth.ParseRefToken(input.RefreshToken)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newAccessToken, err := h.services.Auth.GenerateAccToken(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": newAccessToken,
	})
}
