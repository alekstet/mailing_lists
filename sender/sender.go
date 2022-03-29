package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	url        = "http://localhost:8080/msg"
	timeLayout = "Jan 2, 2006 15:04:05 MST"
)

type Mail struct {
	Id         int
	Text       string
	StartTime  string
	FinishTime string
	Tag        string
	MobileCode int
}

type Client struct {
	Id         int
	Phone      int
	Text       string
	Tag        string
	MobileCode int
}

type Message struct {
	MessageId  int
	Text       string
	SendTime   string
	FinishTime string
	MailId     int
	ClientId   int
}

func MakeSend(send func(c Client), c Mail) error {
	clients := c.Clients
	stopSend := make(chan struct{})
	start, err := time.Parse(timeLayout, c.StartTime)
	if err != nil {
		return err
	}
	finish, err := time.Parse(timeLayout, c.FinishTime)
	if err != nil {
		return err
	}
	delayBeforeStart := start.Sub(time.Now())
	sendDuration := finish.Sub(start)

	go func() {
		<-time.After(sendDuration)
		stopSend <- struct{}{}
	}()

	time.Sleep(delayBeforeStart)

	go func() {
		for _, c := range clients {
			go send(c)
		}
		<-stopSend
		fmt.Println("Fin")
	}()
	return nil
}

func Send(c Client) {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Error with Marshal: %v\n", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		log.Fatalf("Error with request: %v\n", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error with make request: %v\n", err)
	}

	fmt.Println(res.StatusCode)
}

func main() {
	c1 := Client{1, 123, "hello"}
	c2 := Client{2, 126, "hello"}
	mail1 := Mail{1, "hello", "Mar 29, 2022 20:43:00 MSK", "Mar 29, 2022 20:43:20 MSK", []Client{c1, c2}}

	err := MakeSend(Send, mail1)
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	time.Sleep(time.Hour * 24)
}
