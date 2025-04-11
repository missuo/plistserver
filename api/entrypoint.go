package api

import (
	"net/http"

	"github.com/missuo/plistserver/service"

	"github.com/gin-gonic/gin"
)

var (
	cfg *service.Config
	app *gin.Engine
)

func init() {
	cfg = service.InitConfig()
	app = service.Router(cfg)
	gin.SetMode(gin.ReleaseMode)
}

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

var _ http.HandlerFunc = Entrypoint
