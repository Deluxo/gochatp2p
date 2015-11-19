package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type message struct {
	Body string
	Time time.Time
}

func getArg(arg string) string {
	var a string
	for k, v := range os.Args {
		if v == arg {
			a = os.Args[k+1]
			break
		}
	}
	return a
}

func createServer(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			buffer := make([]byte, 2048*2)
			bytesRead, _ := conn.Read(buffer)
			request := string(buffer[0:bytesRead])
			fmt.Println(request)
			conn.Close()
		}
	}
}

func connectServer(addr string, nick string) {
	for {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Millisecond * 1000)
		} else {
			message := createMessage()
			fmt.Fprintf(conn, "["+message.Time.Format("Jan 2 15:04")+"] "+nick+": "+message.Body)
		}
	}
}

func createMessage() message {
	scn := bufio.NewScanner(os.Stdin)
	var body string
	for scn.Scan() {
		line := scn.Text()
		if line[0] == '\x1D' {
			break
		}
		body += line + "\n"
	}
	return message{Body: body, Time: time.Now()}
}

func sendMessage(m message, chatChannel chan message) {
	chatChannel <- m
}

func getMessage(chatChannel chan message) {
	m := <-chatChannel
	fmt.Println(m)
}

func main() {
	nick := getArg("-u")
	for {
		go createServer(getArg("-i"))
		go connectServer(getArg("-o"), nick)
	}
}
