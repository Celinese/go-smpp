package Controller

import (
	"log"
	"os"
	"strings"


	"github.com/fiorix/go-smpp/smpp"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield
	"github.com/fiorix/go-smpp/smpp/pdu/dutext"
	"github.com/fiorix/go-smpp/v2/modl"
	"github.com/gin-contrib/sesions"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/ialects/mysql"
	

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
		PhoneNumber string `form:"phone_number" binding:"required,min=10,max=11"`

		Message string `form:"message" binding:"required,min=1,max=160"`
	}
	var request Request
	if len(request.PhoneNumber) == 10 {
		request.PhoneNumber = "+66" + request.PhoneNumber
	}
	request.PhoneNumber = strings.Replace(request.PhoneNumber, "0", "+66", 1)
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Prepare the SMS message
	msg := &smpp.ShortMessage{
		Src:      request.Sender,
		Dst:      request.PhoneNumber,
		Text:     pdutext.UCS2(request.Message),
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
	log.Println("Sender ID:", request.PhoneNumber)
	log.Println("Sender ID:", request.Message)
	log.Println("Sender ID:", userId)
	log.Println("Sender ID:", userName)

	// TODO: Function Insert the data into the database
	// Get session UserId & Name

	err = model.Submitfrominsert(request.Sender, request.PhoneNumber, request.Message, userName, userId)
	if err != nil {
		log.Fatal(err)
		og.Fatal(err)
}

	// ![Func CeateFile Custom.log & InsertLog To File] //
	// open file and create if non-existent
	file, err := osOpenFile("LogFile/Send.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		og.Fatal(err)
	}
defer file.Close()

	logger := log.New(file, "Send Log", log.LstdFlags)
logger.Println("[Message Sennd] Sender:", request.Sender, "Receiver:", request.PhoneNumber, "Message:", request.Message, "Username:", userName, "UserId:", userId)

	/ ![End Func Intert Log to File] //
}
