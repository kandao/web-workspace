package main

import (
	"os"
	"web-workspace/pkg/svc"
	"web-workspace/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func main() {
	p := utils.NewConfigProperies()
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	slog.Info(utils.ConfigPropertiesToString(p))
	r := gin.Default()
	r.GET("/ping", svc.PingPong)
	r.POST("/example", svc.FiledService)
	// r.PUT("/put_example", func(c *gin.Context) {

	// })
	r.Run(p.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
