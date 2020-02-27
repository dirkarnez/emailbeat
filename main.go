package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/dirkarnez/smail"
)

var (
	email    string
	password string
	smtp     string
	port     int
	interval int
)

const (
	timeFormat = "2006.01.02 15:04:05"
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

	s, err := smail.Dial(
		email,
		password,
		smtp,
		port,
	)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	t := time.NewTicker(time.Duration(time.Second * time.Duration(interval)))
	defer t.Stop()

	for {
		<-t.C

		err = smail.Send(s, fmt.Sprintf("EmailBeat [%s]", time.Now().Format(timeFormat)), "")
		if err != nil {
			panic(err)
		}

		fmt.Println("should have sent")
	}
}
