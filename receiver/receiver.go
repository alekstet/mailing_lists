package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Client struct {
	Id    int
	Phone int
	Text  string
}

func Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	time.Sleep(time.Second * 30)
	data := &Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	json.Unmarshal(body, &data)
	fmt.Println(data.Id, data.Phone, data.Text)
}

func main() {
	router := httprouter.New()
	router.POST("/msg", Test)

	log.Fatal(http.ListenAndServe(":8080", router))
}
