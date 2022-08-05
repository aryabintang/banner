package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerCreat(c *gin.Context, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application")

}
