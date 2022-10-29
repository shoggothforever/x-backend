package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trygorm/app/responses"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	ctx := &gin.Context{}
	responses.SendResponse(ctx, 200, "Pong!", "")
}
