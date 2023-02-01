package Controller

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/fiorix/go-smpp/smpp"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/fiorix/go-smpp/smpp/pdu/pdutext"
	"github.com/fiorix/go-smpp/v2/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// ============================================================== //
// ============================================================== //
//           TODO: Function Controller SEND SMS [POST]            //
// ============================================================== //
// ============================================================== //

func Submitfrom(c *gin.Context) {
	session := sessions.Default(c)

	// TODO: Function Check for session variable

	userId := session.Get("userId").(uint)
	userName := session.Get("userName").(string)

	// ! [ LOAD ENV SETTINGS FILE] //
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file")
	}
	ip := os.Getenv("SMPP_ADDRESS")
	us := os.Getenv("SMPP_USER")
	pw := os.Getenv("SMPP_PASS")

	// TODO: Connect To Server Gateway SMPP 3.4
	tx := &smpp.Transmitter{
		Addr:   ip,
		User:   us,
		Passwd: pw,
	}

	// TODO: Function Bind Status Connect SMPP [Check Error Connection]
	conn := <-tx.Bind()

	if conn.Status() != smpp.Connected {
		log.Fatal(conn.Error())
	}

	// TODO: Check Vaildate From Number & Message [ Number = Max 160 ] [ Phone Number = 11 ]
	// Validate phone number and message
	type Request struct {
		Sender string `form:"sender" binding:"required,min=1,max=50"`
		//	PhoneNumber string `form:"phone_number" binding:"required,len=11"`
		PhoneTo string `form:"Phone_to" binding:"required,min=10,max=11"`

		MessageTo string `form:"Message_to" binding:"required,min=1,max=160"`
	}
	var request Request
	if len(request.PhoneTo) == 10 {
		request.PhoneTo = "66" + request.PhoneTo
	}
	request.PhoneTo = strings.Replace(request.PhoneTo, "0", "66", 1)
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Prepare the SMS message
	msg := &smpp.ShortMessage{
		Src:      request.Sender,
		Dst:      request.PhoneTo,
		Text:     pdutext.UCS2(request.MessageTo),
		Register: pdufield.NoDeliveryReceipt,
	}

	// Send the SMS message
	_, err := tx.Submit(msg)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send SMS message"})
		log.Println(err)
		return
	}

	// Return success message to the client
	c.Redirect(302, "/home")
	/* c.JSON(302, gin.H{"message": "SMS message sent successfully ;( "}) */

	log.Println("Sender ID:", request.Sender)
	log.Println("Sender ID:", request.PhoneTo)
	log.Println("Sender ID:", request.MessageTo)
	log.Println("Sender ID:", userId)
	log.Println("Sender ID:", userName)

	// TODO: Function Insert the data into the database
	// Get session UserId & Name

	err = model.Submitfrominsert(request.Sender, request.PhoneTo, request.MessageTo, userName, userId)
	if err != nil {
		log.Fatal(err)
	}

	// ![Func CeateFile Custom.log & Insert Log To File] //
	currentDate := time.Now().Format("2006-01-02")
	// Create the log file with the date as the file name
	logFile, err := os.OpenFile("LogFile/Send-"+currentDate+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println("[Message Sennd] Sender:", request.Sender, "Receiver:", request.PhoneTo, "Message:", request.MessageTo, "Username:", userName, "UserId:", userId)

	// ![End Func Intert Log to File] //
}
