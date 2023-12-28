package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWithCompany(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("company_id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.GetWithCompany(ctx, id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, res)
}
