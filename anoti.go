package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

func send_mail(username, password, hostname string, port int,
	sender string, receipients []string,
	subject, message string) {
	auth := smtp.PlainAuth("", username, password, hostname)
	msg := "To: "
	for _, r := range receipients {
		msg += r + ", "
	}
	msg = fmt.Sprintf("%s\r\nSubject: %s\r\n\r\n%s\r\n",
		msg, subject, message)
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", hostname, port),
		auth, sender, receipients, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args
	fmt.Println("args: ", args)
	for i, a := range args {
		fmt.Println(i, " ", a)
	}
	port, e := strconv.Atoi(args[4])
	if e != nil {
		log.Fatal("wrong argument to port number!", e)
	}
	send_mail(args[1], args[2], args[3], port, args[5],
		[]string{args[6]}, args[7], args[8])
}
