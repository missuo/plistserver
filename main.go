// main.go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/missuo/plistserver/service"
)

func main() {
	cfg := service.InitConfig()
	fmt.Printf("Plist Generator has been successfully launched! Listening on %v:%v\n", cfg.IP, cfg.Port)

	// Setting the application to release mode
	gin.SetMode(gin.ReleaseMode)
	app := service.Router(cfg)
	app.Run(fmt.Sprintf("%v:%v", cfg.IP, cfg.Port))
}
