package main

import (
	"flag"
	"fmt"
	"gopkg.in/gomail.v2"
	"time"
)

var (
	email string
	password string
	smtp string
	port int
	interval  int
)

const (
	TIME_FORMAT = "2006.01.02 15:04:05"
)

func main() {
	flag.StringVar(&email, "email", "", "Sender and receiver email")
	flag.StringVar(&password, "password", "", "Password for the email")
	flag.StringVar(&smtp, "smtp", "", "Smtp host")
	flag.IntVar(&port, "port", 587, "Smtp port, default 587")
	flag.IntVar(&interval, "interval", 1800, "Heartbeat interval in seconds, default 1800")
	flag.Parse()

	if len(email) < 1 || len(password) < 1 || len(smtp) < 1 {
		fmt.Println("Insufficient arguments")
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", fmt.Sprintf("EmailBeat [%s]", time.Now().Format(TIME_FORMAT)))

	d := gomail.NewDialer(smtp, port, email, password)

	s, err := d.Dial()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	t := time.NewTicker(time.Duration(time.Second * 3))
	defer t.Stop()

	for {
		<- t.C

		err = gomail.Send(s, m)

		if err != nil {
			panic(err)
		}

		fmt.Println("should have sent")
	}
}