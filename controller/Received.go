package Controller

import (
	"net/http"

	"github.com/fiorix/go-smpp/v2/model"
	"github.com/gin-gonic/gin"
)

/* var err error */

//GetUsers ... Get all users
func GetSms(c *gin.Context) {
	var received []model.Received
	err := model.GetAllSMS(&received)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, received)
	}
}
