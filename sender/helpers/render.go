package helpers

import (
	"encoding/json"
	"net/http"
)

type error_resp struct {
	Error string `json:"error"`
}

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	resp := error_resp{Error: err.Error()}
	Render(w, r, code, resp)
}

func Render(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	if data == nil {
		w.WriteHeader(code)
		return
	}
	if data != nil {
		resp, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(code)
		w.Write(resp)
	}
}
