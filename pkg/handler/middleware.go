package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "UserId"
	userList
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusBadRequest, "empty auth head")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusBadRequest, "invalid auth header")
		return
	}
	if len(headerParts[1]) == 0 {
		NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}
	userId, err := h.services.Auth.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}
func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return "", errors.New("user id not found")
	}
	idString, ok := id.(string)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is invalid type")
		return "", errors.New("user id not found")
	}
	return idString, nil
}
