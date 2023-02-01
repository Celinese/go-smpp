package model

import (
	"time"
)

type Received struct {
	Id               uint   `json:"id"`
	Sender           string `json:"Sender"`
	Phone_customer   string `json:"Phone_customer"`
	Message_customer string `json:"Message_customer"`
	Date_received    time.Time
}

type Sendsms struct {
	Id          uint   `json:"id"`
	Sender      string `json:"sender"`
	Phone_to    string `json:"Phone_to"`
	Message_to  string `json:"Message_to"`
	User_name   string `json:"User_name"`
	User_id     string `json:"User_id"`
	hand_on     string `json:"hand_on"`
	Date_insert time.Time
}

type User struct {
	Id       uint   `json:"id"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
	Role     string `json:"Role"`
	Phone    string `json:"Phone"`
	Date     time.Time
}

func (b *Received) TableName() string {
	return "received"

}
func (b *Sendsms) TableName() string {
	return "Sendsms"

}
