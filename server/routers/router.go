package routers

import (
	v1 "github.com/athagi/hello-copilot/server/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	// r.POST("/auth", api.GetAuth)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	{

		apiv1.GET("/healthz", v1.GetHealthz)
		apiv1.GET("/tags", v1.GetTags)
		apiv1.GET("/images", v1.ListImages)
		apiv1.POST("/images/upload", v1.UploadImage)
		apiv1.GET("/images/download", v1.DownloadImage)
		apiv1.POST("/images/delete", v1.DeleteImage)
		return r
	}

}
