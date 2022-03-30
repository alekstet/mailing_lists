package process

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/mailing_lists/sender/api/models"
	"github.com/alekstet/mailing_lists/sender/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) AddClient(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var exists int
	data := &models.Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}
	json.Unmarshal(body, &data)

	err = s.Db.QueryRow(
		`SELECT EXISTS
		(SELECT * FROM clients 
			WHERE id = $1)`, data.Id).Scan(&exists)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}

	if exists == 1 {
		helpers.Error(w, r, 400, errors.New("client already exists"))
		return
	} else {
		_, err = s.Db.Exec(
			`INSERT INTO clients (id, phone, mobile_code, tag) 
			VALUES ($1, $2, $3, $4)`, data.Id, data.Phone, data.MobileCode, data.Tag)
		if err != nil {
			helpers.Error(w, r, 500, err)
			return
		}
	}

	helpers.Render(w, r, 201, nil)
}
