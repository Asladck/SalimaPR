package handler

import (
	"SalimProject/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) generateOutfit(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	id, err := h.services.Outfit.GenerateOutfit(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllOutfits(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	outfits, err := h.services.Outfit.GetAllOutfits(userId)

	c.JSON(http.StatusOK, dto.GetAllOutFitsResponse{
		Data: outfits,
	})
}
func (h *Handler) getOutfitById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	outId := c.Param("id")
	outfit, err := h.services.Outfit.GetOutfitById(userId, outId)
	c.JSON(http.StatusOK, dto.GetOutFitResponse{
		Data: outfit,
	})
}
func (h *Handler) deleteOutfitById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	outId := c.Param("id")
	err = h.services.Outfit.DeleteOutfitById(userId, outId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
