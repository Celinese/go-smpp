package Controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/fiorix/go-smpp/v2/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDb() error {
	var err error
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	return nil
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func HandleLogin(c *gin.Context) {
	if err := InitDb(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	var user model.User
	email := c.PostForm("email")
	password := c.PostForm("password")

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to the database"})
		return
	}
	if user.Password != password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Create session
	session := sessions.Default(c)
	session.Set("userId", user.Id)
	session.Set("userName", user.Name)
	session.Set("userRole", user.Role)
	session.Set("userPhone", user.Phone)
	session.Save()
	c.Redirect(http.StatusFound, "/home")

}

/* ==================[HOME]========================= */

func ShowHomePage(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userId")
	userName := session.Get("userName")
	userRole := session.Get("userRole")
	userPhone := session.Get("userPhone")

	if userId == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	var user model.User
	if err := db.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to the database"})
		return
	}

	t, err := template.ParseFiles("templates/header.html", "templates/home.html", "templates/footer.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing template files"})
		return
	}

	// TODO:Execute the template Headers.HTML

	err = t.ExecuteTemplate(c.Writer, "header.html", gin.H{
		"userId":    userId,
		"userName":  userName,
		"userRole":  userRole,
		"userPhone": userPhone,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error executing template Headers: " + err.Error()})
		return
	}

	// TODO:Execute the template HomePage.HTML

	err = t.ExecuteTemplate(c.Writer, "home.html", gin.H{
		"userId":    userId,
		"userName":  userName,
		"userRole":  userRole,
		"userPhone": userPhone,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error executing template HomePage: " + err.Error()})
		return
	}

	// TODO:Execute the template Footer.HTML

	err = t.ExecuteTemplate(c.Writer, "footer.html", gin.H{
		"userId":    userId,
		"userName":  userName,
		"userRole":  userRole,
		"userPhone": userPhone,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error executing template Footer: " + err.Error()})
		return
	}

}

func ShowLogsPage(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userId")
	userName := session.Get("userName")
	userRole := session.Get("userRole")
	userPhone := session.Get("userPhone")

	if userId == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	var user model.User
	if err := db.First(&user, userId).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// ! [START]
	/* userIdValue, ok := userId.(uint)
	if !ok {
		// Handle the error
		log.Println("Invalid type for userId")
	}
	sendsms, err := model.QuerySendLogs(userIdValue)
	if err != nil {
		// Handle the error
		log.Println(err)
	} */
	// ! [END]
	t, err := template.ParseFiles("templates/header.html", "templates/Logs.html", "templates/footer.html")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(c.Writer, "header.html", gin.H{
		"userId":    userId,
		"userName":  userName,
		"userRole":  userRole,
		"userPhone": userPhone,
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = t.ExecuteTemplate(c.Writer, "Logs.html", gin.H{
		"userId":    userId,
		"userName":  userName,
		"userRole":  userRole,
		"userPhone": userPhone,
		/* "sendsms":   sendsms, */
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = t.ExecuteTemplate(c.Writer, "footer.html", gin.H{
		"userId":    userId,
		"userName":  userName,
		"userRole":  userRole,
		"userPhone": userPhone,
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

}

/* !!==================[ShowLogsReceived]=========================!! */

func ShowLogsReceived(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userId")
	userName := session.Get("userName")
	userRole := session.Get("userRole")
	userPhone := session.Get("userPhone")

	if userId == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	var user model.User
	if err := db.First(&user, userId).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// ! Connect Function Query Database ! //
	/* receivedsms, err := model.ShowLogsReceived()
	if err != nil {
		// Handle the error
		log.Println(err)
	} */
	// ! END Connect Function Query Database ! //

	t, err := template.ParseFiles("templates/header.html", "templates/Logs.html", "templates/footer.html")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(c.Writer, "header.html", gin.H{
		"userId":    userId,
		"userName":  userName,
		"userRole":  userRole,
		"userPhone": userPhone,
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = t.ExecuteTemplate(c.Writer, "LogsReceived.html", gin.H{
		"userId":    userId,
		"userName":  userName,
		"userRole":  userRole,
		"userPhone": userPhone,
		/* "sendsms":   receivedsms, */
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = t.ExecuteTemplate(c.Writer, "footer.html", gin.H{
		"userId":    userId,
		"userName":  userName,
		"userRole":  userRole,
		"userPhone": userPhone,
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

}

/* ===== [ Fuction LOGOUT] ======= */

func HandleLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}
