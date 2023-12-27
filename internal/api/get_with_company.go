package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWithCompany(ctx *gin.Context) {
	NewErrorResponse(ctx, http.StatusNotImplemented, "method not implemented")
}
