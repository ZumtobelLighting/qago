package main

import (
	"log"
	"net/smtp"
)

func main() {
	send("hello there Sapna")
}

func send(body string) {
	from := "dlqa02@gmail.com"
	pass := "welcome2DL"
	to := "sparikh@digitallumens.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: testing sending email from golang\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
