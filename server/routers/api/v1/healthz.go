package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func GetHealthz(c *gin.Context) {

	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": "server is alive",
	})

}
