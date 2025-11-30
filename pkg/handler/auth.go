package handler

import (
	"SalimProject/pkg/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input dto.SignUpInput
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

func (h *Handler) signIn(c *gin.Context) {
	var input dto.SignInInput

	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	access, refresh, err := h.services.Auth.SignIn(input)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func (h *Handler) refreshHandler(c *gin.Context) {
	var input dto.RefreshInput
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.
			StatusBadRequest, err.Error())
		return
	}
	userId, err := h.services.Auth.ParseRefresh(input.RefreshToken)
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
