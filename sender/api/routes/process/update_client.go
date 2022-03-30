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

func (s *S) UpdateClient(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	data := &models.Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}
	json.Unmarshal(body, &data)

	row, err := s.Db.Exec(
		`UPDATE users_questions SET Answer = $1, Updated_at = $2
		WHERE Question_id = $3 AND User_nickname = $4`,
		data.Phone, data.MobileCode, data.Tag)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}
	isUpdated, _ := row.RowsAffected()
	if isUpdated == 0 {
		helpers.Error(w, r, 500, err)
		return
	}

	helpers.Render(w, r, 200, nil)
}
