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
	Clients    []Client
}

type Client struct {
	Id    int
	Phone int
	Text  string
}

func MakeSend(f func(c Client), c Mail) error {
	stopSend := make(chan struct{})
	start, err := time.Parse(timeLayout, c.StartTime)
	if err != nil {
		return err
	}
	finish, err := time.Parse(timeLayout, c.FinishTime)
	if err != nil {
		return err
	}
	duration := start.Sub(time.Now())

	go func() {
		go func() {
			for {
				needToStop := finish.Sub(time.Now())
				if needToStop.Seconds() < 0.0 {
					fmt.Println("Time!")
					stopSend <- struct{}{}
					break
				}
			}
		}()
		time.Sleep(duration)

		clients := c.Clients
		go func() {
			for _, c := range clients {
				go f(c)
			}
			for {
				<-stopSend
				fmt.Println("Fin")
				return
			}
		}()

	}()

	time.Sleep(time.Hour * 24)

	return nil
}

func Send(c Client) {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Error with Marshal: %v\n", err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		log.Fatalf("Error with request: %v\n", err)
	}

	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error with make request: %v\n", err)
	}

	fmt.Println(res.StatusCode)
}

func main() {
	c1 := Client{1, 123, "hello"}
	c2 := Client{2, 126, "hello"}
	new := Mail{1, "hello", "Mar 29, 2022 14:46:00 MSK", "Mar 29, 2022 14:46:20 MSK", []Client{c1, c2}}

	err := MakeSend(Send, new)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	time.Sleep(time.Hour * 24)
}
