##### [ปัญหา / บัค ที่พบ] #####
## [16/1/2566] ##
- อัพเดตเมื่อเสาร์อาทิตย์ที่ผ่านมา ได้ทำการรวม ตัวรับและส่งไปได้ประมาณ 60 % แล้ว อุปสรรค์ ในการร่วม คือตัวแปร หลายๆตัวไม่เหมือนกัน ฟังชั่นการเขียนคนละแบบกันเนื่องจากผมเอามาจากคนละโปรเจค จึงต้องแก้เยอะ บางฟังชั่นไม่มี ต้องเอาไปแก้ให้เข้ากับตัวแปรหลักเปิด Doc โปรเจคหลักดู ฟังชั่นแต่ละตัวที่เขาทำไว้ บางที รันผ่านไม่มีเออเร่ออะไร แต่ พอเรียกฟังชั่น POST ของ API กับไม่ทำงาน สถานะ 
Sending request... ใน POST MAN 
## [13/1/2566] ##
- ฝั่ง Receiver พบปัญหาในส่วนของ API เวลาเรา GET ข้อมูลปกติ (ในกรณีที่ยังไม่มี SMS ตอบกลับมา) จะสามารถใช้งานได้ปกติ GET ข้อมูลได้ แต่ถ้ามีข้อความมา จะไม่สามารถ GET Data ในส่วนของ API ได้เลย API Error : 404 Not found 
( หาข้อมูลมาน่าจะเป็นที่ส่วนของ SQL ไม่สามารถ Connect ได้ เนื่องจากติด  Loop รอรับข้อมูล มันทำงานอยู่ มันเลยติดอยู่ตรงนั้น เพราะฟังชั่นการทำงาน จะแบ่งเป็น 2 ส่วน 
- ส่วนแรก จะเป็นการเชื่อมต่อ Config ของ DB พอ Connect ได้แล้วจะ ทำการ [defer Config.DB.Close()] จาก Doc เมื่อใช้งานเสร็จแล้วจะต้องทำการปิด เพื่อออกจากลูปนั้นออก แล้วไปทำงานในสเต๊ปถัดไป
- ส่วนที่สอง จะเป็นการเรียกฟังชั่น เชื่อมต่อและ Config ตัว SMPP Receiver ในส่วนนี้จะมีการ Setting Timer for Received SMS Message ไว้
	    // TODO: After Function หลังจากการเชื่อมต่อสำเร็จจะนับคลูดาว์นถอยหลัง ตามเวลาที่ตั้งไว้ หน่วยเป็น วินาที
	    // TODO: [3600 = 1 Hr.] & [5400 = 1.30 hr.] & [7200 = 2 hr.] & [9200 = 2.30 hr.] & [10800 = 3 hr] & [12600 = 3.30 hr] & [14400 = 4 hr.]

	    // TODO: หลักการทำงาน เงื่อนไขแรก ถ้าโปรแกรมรันไปแล้ว จะนับเวลาถอยหลังตามที่เราตั้งไว้ ถ้าหากไม่มีข้อความส่งมา ภายในเวลา(ในที่นี้ตั้งไว้ 1 ชั่วโมง) ก็จะจบการทำงาน สามารถปรับตั้งเวลาได้
	        time.AfterFunc(3600*time.Second, func() {

	    r.Close()
	})

พอครบเวลาที่กำหนด ถึงจะ ปิดการทำงานของฟังชั่นนี้ จึงทำให้เวลามี ข้อความเริ่มเข้ามาข้อความแรก จะไม่สามารถ วนไปใช้งานฟังชั่น GET API ได้เลย เนื่องจากไม่สามารถ Connect Database ได้


## [สิ่งที่สงสัย] ##
- ตัวส่งใช้การ Login Username password แล้วพอกดส่ง จะเก็บ LOG username หรือ id เวลาเรา get เราก็จะ getตาม username ,id 
- ตัวรับ ข้อมูล จะถูกส่งกลับมาในเซิฟเวอร์ เราจะเช็คจากอะไรในการที่จะ get ข้อมูลออกมาว่าข้อความที่ส่งกลับมาเป็นของ เจ้าไหน (ในกรณีUser ใช้งาน) [ ทราบวิธีการแล้ว ไปเทียบโค้ดเก่าของPHP ใช้ฟังชั่น JOIN TABLE DB มาเทียบเบอร์โทรศัพท์ผู้เบอร์ลูกค้า]

## [12/1/2566] ##
- ฝั่ง Receiver รับ Message มาแล้วทำการ encode ออกมาเป็น gsm7 ซึ่งไม่รองรับภาษาไทย ต้องแก้ไขให้เป็น ucs2 = utf-16 :white_check_mark: [Sucessfully Updated 12/1/2566]




##### [ แก้ไขเรียบร้อยแล้ว ] #####


## [12/1/2566] ##
- แก้ไขการ decode เป็น UTF เรียบร้อยแล้ว วิธีคือ เขียนรับค่า Message มาเก็บในรูปแบบ String แล้วใช้ Library golang.org/x/text/encoding/unicode ในการdecode ให้เป็น utf 

## [11/1/2566] ##
- เพิ่ม Config ให้สามารถเรียกใช้งาน ภายใต้ .env file [MYSQL STINGS , SMPP STINGS, SECRET_API]
- เขียน Auto Migrate SQL สร้างตาราง ออโต้
- เขียน MVC Controller , model , Config ไว้รอเรียบร้อยแล้ว
- แก้ไข ปัญหาการ Insert แล้ว data ไม่เข้า (ปัญหาคือ ต้องเอาค่าที่ PDU.SMPP รับมา มาเก็บเข้าในรูปแบบ String ก่อน และทำการนำมาแยก แพคออก เพราะข้อมูลที่ส่ง มาเป็นอารเรย์ )
- แก้ไข บัคจาก For loop ในการวนการทำงานของตัวตั้งเวลา รับข้อความแล้วทำให้ เวลารอรับข้อความแล้ว insert ข้อมูล error เข้ามารัวๆ (แก้ไขเอาตัว ฟังชั่น Insert ไปนอกลูป แล้วทดสอบไม่มีปัญหา เรื่องส่งข้อความแล้วไม่เข้า หรือ ส่งได้รอบเดียว (ในส่วนนี้ต้องรอเทสต่อๆไปว่าจะมีปัญหาเพิ่มไหม)) 
- ตัวรับ Received เชื่อมต่อ Mysql  เก็บ Log = [Phone-Message-Sender-Date&Time] , Logfile = [Receiver.log] เรียบร้อยแล้ว

## [09/1/2566] ##
- แก้ไขเรื่อง Time Seconds เรียบร้อย เขียน LogTerminal แสดง Count Time ที่โปรแกรมรันไปว่ากี่นาทีแล้วจึงหยุดทำงาน
- เขียน Logfile เก็บ Log Recei[ เบอร์ผู้รับ & เบอร์ผู้ส่ง & ข้อความ ] , กรณีที่ Error รับข้อความมาไม่ได้หรือมีปัญหาการเชื่อมต่อ ก็จะ   
  Save ลง Logfile ไวเหมือนกัน ข้อความใน Log ที่ Save มา [เบอร์ผู้ส่ง & MessageId=เช่น 23010913020466977063999 & วันที่ & โค้ดเออเร่อ]
- ตัว Recevied เชื่อม DB ไว้เรียบร้อย ทำ Restful API ไว้ GET ข้อความเรียบร้อย 


##### [ สิ่งที่ต้องทำต่อ ] #####
- ทำการรวมโปรเจคงานตัวรับและตัวส่ง
- ตัวส่งทำ Author [X-API-Key],[Username-Password-Token]
- ตัวส่งเชื่อมต่อ Mysql เก็บ Log = [Phone-Message-Sender-Date&Time] , Logfile = [Sender.log]

- สร้างตัว API Receiver เพื่อที่จะทำให้สามารถเรียกใช้งานภายนอกได้ กรณีต้องการดึงข้อมูลจาก DB พร้อมกับทำ Author 
  token-key-username-password [ในส่วนของAPI เรียกใช้งานข้อมูลทำไว้แล้วเหลือทำตัว Auth พร้อมกับตัวส่ง]
- การ Auth Token จะทำได้ต่อจากการที่รวมโปรเจคสำเร็จแล้ว เนื่องจากจะทำรวบเป็นโฟเดอร์ให้เป็นสัดส่วน



##### [ Design Create Database ] #####

[Table Received]
Id = AUTO_INCREMENT not null
Phone = varchar(10) not null
Message = text 
Dtime = datetime not null


[Table Sender]
Id = AUTO_INCREMENT not null
PhoneTo = varchar(10) not null
PhoneFrom = varchar(10) not null
Message = text 
Dtime = datetime not null




##### [Workflows Project]  #####
go-smpp/
|
| - api/
|   | - server.go/          # RUN: API

| - authorization/
|   | - auth.go                 # Authorization Token -------------------------------- [ยังไม่เสร็จรอขั้นตอนเกือบสุดท้ายเผื่อเจอปัญหา]

| - config/
|   | - config.go                 # config for db connection

| - controller/
|   | - todo.controllers.go       # func to handling data

  - middlewares/
|   | - middlewares.go       # HandlerFunc / Herder Token / JWT ------------------------ [ยังไม่เสร็จรอขั้นตอนเกือบสุดท้ายเผื่อเจอปัญหา]

| - model/
|   | - user.go            # fucntion for model user GET / POST / DELETE /
      - UserModel.go       # function Create Datamodel Insert To Database

| - routes/
|   | - routes.go                 # handling HTTP methods

| - smpp/                         
       - encoding/
           - gsm7.go      # Libary endcode Text SMS ------ [ต้องหาไฟล์ที่เรียกใช้งาน UTF-8 / UTF-16 เนื่องจาก gsm7 ไม่สามารถอ่านภาษาไทยได้]
           - gsm7_test.go

  - pdu/                         
       - pdufield/              # Libary PDU Field SMPP

       - pdutext/               # Libary PDU TEXT SMPP Convert

       - pdutlv/                # Libary PDU TLV [SMPP]
    
  - routes/    
|   | - routes.go                 # handling HTTP methods

| - .env                              # environtment vars
| - go.mod                            # golang modules
| - main.go                           # golang main data