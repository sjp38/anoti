package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

type task struct {
	Cmd         string
	FailMsg     string
	Receipients []string
}

var tasks []task

func read_tasks(filePath string) {
	c, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("failed to read task file %s: %s\n", filePath, err)
		return
	}

	if err := json.Unmarshal(c, &tasks); err != nil {
		log.Printf("failed to unmarshal task file %s: %s\n",
			filePath, err)
		return
	}

	for i, t := range tasks {
		fmt.Printf("task %d\n", i)
		fmt.Printf("cmd: %s\nFail msg: %s\nReceipients:%s\n\n",
			t.Cmd, t.FailMsg, t.Receipients)
	}
}

func write_sample_tasks(filePath string) {
	tasks = []task{task{Cmd: "abc_cmd", FailMsg: "fail_msg",
		Receipients: []string{"receipients 1", "1"}}}
	bytes, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("failed to marshal tasks: %s\n", err)
		return
	}

	if err := ioutil.WriteFile(filePath, bytes, 0600); err != nil {
		log.Printf("failed to write marshaled tasks: %s\n", err)
		return
	}
}

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
	if len(args) == 2 {
		read_tasks(args[1])
		return
	}

	port, e := strconv.Atoi(args[4])
	if e != nil {
		log.Fatal("wrong argument to port number!", e)
	}
	send_mail(args[1], args[2], args[3], port, args[5],
		[]string{args[6]}, args[7], args[8])
}
