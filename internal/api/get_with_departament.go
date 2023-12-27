package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWithDepartament(ctx *gin.Context) {
	NewErrorResponse(ctx, http.StatusNotImplemented, "method not implemented")
}
