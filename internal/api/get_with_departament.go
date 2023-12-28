package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWithDepartament(ctx *gin.Context) {

	dep := ctx.Param("dep")
	com, err := strconv.Atoi(ctx.Param("com"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.GetWithDepartament(ctx, dep, com)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}
