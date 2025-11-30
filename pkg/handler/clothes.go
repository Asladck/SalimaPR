package handler

import (
	"SalimProject/models"
	"SalimProject/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addClothes(c *gin.Context) {
	var input dto.AddClothesInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	item := &models.Clothes{
		UserId:   userID,
		Name:     input.Name,
		Category: input.Category,
		Color:    input.Color,
		Season:   input.Season,
		Material: input.Material,
		ImageURL: input.ImageURL,
	}

	id, err := h.services.Clothes.AddClothes(item)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllClothes(c *gin.Context) {
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
	c.JSON(http.StatusOK, dto.GetAllClothResponse{
		Data: items,
	})
}

func (h *Handler) getClothesById(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	itemID := c.Param("id")

	item, err := h.services.Clothes.GetClothesById(userID, itemID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": item,
	})
}

func (h *Handler) deleteClothesById(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	itemID := c.Param("id")
	err = h.services.Clothes.DeleteClothesById(userID, itemID)

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
func (h *Handler) updateClothesById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	itemID := c.Param("id")
	var input dto.ClothesUpdateInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Clothes.UpdateClothesById(userId, itemID, input)
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
