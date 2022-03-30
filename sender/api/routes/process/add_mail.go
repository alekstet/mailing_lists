package process

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/alekstet/mailing_lists/sender/api/models"
	"github.com/alekstet/mailing_lists/sender/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

const (
	url        = "http://localhost:8080/msg"
	timeLayout = "Jan 2, 2006 15:04:05 MST"
)

func MakeSend(send func(c models.Client), mail models.Mail, clients []models.Client) error {
	stopSend := make(chan struct{})
	start, err := time.Parse(timeLayout, mail.StartTime)
	if err != nil {
		return err
	}
	finish, err := time.Parse(timeLayout, mail.EndTime)
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
		for _, j := range clients {
			go send(j)
			<-stopSend
			fmt.Println("Fin")
		}
	}()
	return nil
}

func Send(c models.Client) {
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

func (s *S) AddMail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var exists int
	data := &models.Mail{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}
	json.Unmarshal(body, &data)

	err = s.Db.QueryRow(
		`SELECT EXISTS
		(SELECT * FROM mails 
			WHERE id = $1)`, data.Id).Scan(&exists)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}

	if exists == 1 {
		helpers.Error(w, r, 400, errors.New("mail already exists"))
		return
	} else {
		_, err = s.Db.Exec(
			`INSERT INTO mails (id, start_time, end_time, text, mobile_code, tag) 
			VALUES ($1, $2, $3, $4, $5, $6)`, data.Id, data.StartTime, data.EndTime, data.Text, data.MobileCode, data.Tag)
		if err != nil {
			helpers.Error(w, r, 500, err)
			return
		}
	}

	rows, err := s.Db.Query(
		`SELECT * FROM clients 
		WHERE tag = $1 OR mobile_code = $2`, data.Tag, data.MobileCode)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}
	defer rows.Close()

	clients := []models.Client{}

	for rows.Next() {
		p := models.Client{}
		err := rows.Scan(&p.Id, &p.MobileCode, &p.Tag)
		if err != nil {
			helpers.Error(w, r, 500, err)
			return
		}
		clients = append(clients, p)
	}

	helpers.Render(w, r, 201, nil)

	err = MakeSend(Send, *data, clients)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
