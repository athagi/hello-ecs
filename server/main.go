package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/athagi/hello-copilot/pkg/logging"
	"github.com/athagi/hello-copilot/pkg/setting"
	"github.com/athagi/hello-copilot/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("init")
	setting.Setup()
	// models.Setup()
	logging.Setup()
	// util.Setup()
}

func main() {

	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
	// r := gin.Default()
	// hash := os.Getenv("COMMIT_HASH")
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 		"hash":    hash,
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
