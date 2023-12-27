package api

import (
	"net/http"
	"strconv"

	"github.com/drizzleent/emplyees/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var employe model.Employee

	if err := ctx.BindJSON(&employe); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	employe.Id = id
	err = h.service.Update(ctx, &employe)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data updated": employe,
	})
}
