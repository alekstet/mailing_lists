package testutils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"unsafe"

	"github.com/alekstet/mailing_lists/sender/api/routes"
	"github.com/alekstet/mailing_lists/sender/conf"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Store struct {
	suite.Suite
	Db *sql.DB
}

func New() *conf.Store {
	return &conf.Store{
		Log:    logrus.New(),
		Routes: httprouter.New(),
	}
}

func SendForm(t *testing.T, db *sql.DB, method string, target string, i interface{}) *httptest.ResponseRecorder {
	n := New()
	n.Db = db
	mux := routes.Routes(*n)

	var body io.Reader
	switch v := i.(type) {
	case url.Values:
		body = strings.NewReader(v.Encode())
	case string:
		body = strings.NewReader(v)
	case nil:
	default:
		res, err := json.Marshal(v)
		assert.Nil(t, err)
		body = strings.NewReader(*(*string)(unsafe.Pointer(&res)))
	}

	fmt.Println(body)

	r := httptest.NewRequest(method, target, body)
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	return w
}
