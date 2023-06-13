package main

import (
	"log"

	"gopkg.in/mail.v2"
)

func main() {
	m := mail.NewMessage()
	m.SetHeader("From", "e.cayir2022@gtu.edu.tr")
	m.SetHeader("To", "erdogancayir98@gmail.com")
	m.SetHeader("Subject", "Confirmation code")
	m.SetBody("text/plain", "Your confirmation code is 123456.")

	d := mail.NewDialer("smtp.office365.com", 587, "e.cayir2022@gtu.edu.tr", "boston1907..")
	d.StartTLSPolicy = mail.MandatoryStartTLS

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}
