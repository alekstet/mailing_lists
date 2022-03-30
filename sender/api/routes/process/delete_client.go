package process

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/mailing_lists/sender/api/models"
	"github.com/alekstet/mailing_lists/sender/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) DeleteClient(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	data := &models.Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}
	json.Unmarshal(body, &data)
	_, err = s.Db.Exec(
		`DELETE FROM users_questions 
		WHERE id = $1`, data.Id)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}

	helpers.Render(w, r, 200, nil)
}
