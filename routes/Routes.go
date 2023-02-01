package Routes

import (
	"github.com/fiorix/go-smpp/v2/Controller"
	"github.com/fiorix/go-smpp/v2/middlewares"
	"github.com/fiorix/go-smpp/v2/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	r := gin.Default()
	// apply middleware to specific routes
	//r.POST("/sms", middlewares.BasicAuthMiddleware, middlewares.APIKeyAuthMiddleware, Controller.SendSms)

	// TODO: Group 1   GET Messages from Receive [AND] GET Messages from JOIN 2 TABLES
	connx1 := r.Group("/received")
	{
		connx1.GET("readall", middlewares.BasicAuthMiddleware, middlewares.APIKeyAuthMiddleware, Controller.GetSms)
		connx1.GET("test/:phone", middlewares.BasicAuthMiddleware, middlewares.APIKeyAuthMiddleware, model.GetItemtest)
	}
	// TODO: Group 2  SEND SMS AUTH & KEY-API [AND] GET Messages from SEND
	connx2 := r.Group("/send")
	{
		connx2.GET("sendsms", middlewares.BasicAuthMiddleware, middlewares.APIKeyAuthMiddleware, Controller.GetSMSSend)
		connx2.POST("sms", middlewares.BasicAuthMiddleware, middlewares.APIKeyAuthMiddleware, Controller.SendSms)
	}

	/* connx3 := r.Group("/none")
	{
		connx3.POST("submit", Controller.Submitfrom)
		connx3.POST("sms", Controller.SendSms)
	} */

	/* =================================================[Routes WebLogin]===================================================== */
	// Serve the CSS stylesheet

	r.Static("/css", "templates/assets/")
	r.Static("/js", "templates/assets/")
	r.Static("/img", "templates/images/")
	r.LoadHTMLGlob("templates/*.html")
	store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("mysession", store))
	r.Use(sessions.Sessions("sessionName", store))
	r.GET("/login", Controller.ShowLoginPage)
	r.POST("/login", Controller.HandleLogin)
	r.GET("/home", Controller.ShowHomePage)
	r.GET("/Logs", Controller.ShowLogsPage)
	r.GET("/logout", Controller.HandleLogout)

	r.POST("/submit", Controller.Submitfrom)

	r.Run(":8888")
	return r

}

/* 	// Serve the CSS stylesheet
r.Static("/css", "templates/assets/css/")
r.Static("/js", "templates/assets/js/")
r.Static("/img", "templates/images/")
r.LoadHTMLGlob("templates/*")
store := cookie.NewStore([]byte("secret"))
r.Use(sessions.Sessions("mysession", store))
r.GET("/login", Controller.ShowLoginPage)
r.POST("/login", Controller.HandleLogin)
r.GET("/home", Controller.ShowHomePage)
r.GET("/header", Controller.ShowHomePage)
r.GET("/footer", Controller.ShowHomePage)
*/
