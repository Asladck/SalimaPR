package handler

import (
	"SalimProject/models"
	"SalimProject/pkg/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (h *Handler) addClothes(c *gin.Context) {
	var input dto.AddClothesInput
	if err := c.BindJSON(&input); err != nil {
		logrus.Infof("Invalid input: %v", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userID, err := getUserId(c)
	if err != nil {
		logrus.Info("Can`t find this user")
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	id, err := h.services.Clothes.AddClothes(models.Clothes{
		UserId:    userID,
		Name:      input.Name,
		Category:  input.Category,
		Color:     input.Color,
		Season:    input.Season,
		Material:  input.Material,
		ImageURL:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllItemResponse struct {
	Data []models.Clothes `json:"data"`
}

func (h *Handler) getClothesByUserId(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	items, err := h.services.Clothes.GetClothesByUserId(userID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllItemResponse{
		Data: items,
	})
}

func (h *Handler) deleteClothesById(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	itemID := c.Param("id")
	err = h.services.Clothes.DeleteClothes(itemID)

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
