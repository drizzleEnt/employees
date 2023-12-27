package api

import (
	"net/http"

	"github.com/drizzleent/emplyees/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(ctx *gin.Context) {
	var employe model.Employee

	if err := ctx.BindJSON(&employe); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Create(ctx, &employe)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
