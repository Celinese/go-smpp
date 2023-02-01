package api

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fiorix/go-smpp/smpp"
	"github.com/fiorix/go-smpp/smpp/pdu"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/fiorix/go-smpp/v2/Config"
	"github.com/fiorix/go-smpp/v2/Routes"
	"github.com/fiorix/go-smpp/v2/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang.org/x/text/encoding/unicode"
)

var err error

func Run() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&model.Received{}, &model.Sendsms{}, &model.User{})
	//Config.DB.AutoMigrate()

	// Get the current time
	now := time.Now()
	f := func(p pdu.Body) {
		switch p.Header().ID {
		case pdu.DeliverSMID:
			f := p.Fields()

			// TODO: Function Extract the sender, receiver, and message from the PDU
			address := f[pdufield.SourceAddr].String()
			phone := f[pdufield.DestinationAddr].String()
			message := f[pdufield.ShortMessage].String()

			// ![ Function encode message from byte To UTF] //
			// convert the message from binary to a string in UCS2 encoding
			decoder := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder()
			ucs2Message, _ := decoder.String(string(message))

			log.Printf("Message :> From=%q To=%q Msg=%q Date=%q",
				address, phone, message, now)

			// TODO: Function Insert the data into the database
			err := model.InsertDB(address, phone, ucs2Message)
			if err != nil {
				log.Fatal(err)
			}

			// ![Func CeateFile Custom.log & Insert Log To File] //
			// open file and create if non-existent
			file, err := os.OpenFile("LogFile/received.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			logger := log.New(file, "Received Log", log.LstdFlags)
			logger.Println("[Message Receiver] from:", address, "to:", phone, "Message:", ucs2Message)

			// ![End Func Intert Log to File] //
		}

	}

	// ! [ LOAD ENV SETTINGS FILE] //
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file")
	}
	ip := os.Getenv("SMPP_ADDRESS")
	us := os.Getenv("SMPP_USER")
	pw := os.Getenv("SMPP_PASS")

	// TODO: Connect TO SERVER GETAWAY SMS NT

	r := &smpp.Receiver{
		Addr:    ip,
		User:    us,
		Passwd:  pw,
		Handler: f,
	}
	conn := r.Bind()

	// TODO: After Function หลังจากการเชื่อมต่อสำเร็จจะนับคลูดาว์นถอยหลัง ตามเวลาที่ตั้งไว้ หน่วยเป็น วินาที
	// TODO: [3600 = 1 Hr.] & [5400 = 1.30 hr.] & [7200 = 2 hr.] & [9200 = 2.30 hr.] & [10800 = 3 hr] & [12600 = 3.30 hr] & [14400 = 4 hr.]

	// TODO: หลักการทำงาน เงื่อนไขแรก ถ้าโปรแกรมรันไปแล้ว จะนับเวลาถอยหลังตามที่เราตั้งไว้ ถ้าหากไม่มีข้อความส่งมา ภายในเวลา(ในที่นี้ตั้งไว้ 20 วินาที) ก็จะจบการทำงาน
	time.AfterFunc(3600*time.Second, func() {

		r.Close()
	})

	Routes.SetupRouter()

	// TODO: Print connection status (Connected, Disconnected, etc).
	for c := range conn {
		log.Println("SMPP connection status:", c.Status())

	}

}
