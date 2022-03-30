package process_test

import (
	"testing"
	"time"

	"github.com/alekstet/mailing_lists/sender/api/models"
	"github.com/alekstet/mailing_lists/sender/testutils"

	"github.com/stretchr/testify/suite"
)

type DeleteAnswerSuite struct {
	testutils.Store
}

func (d *DeleteAnswerSuite) SetupSuiteDeleteAnswer() {
	d.Db = testutils.LoadDatabase()
}

func (d *DeleteAnswerSuite) SignUp() {
	var r models.SignUp
	r.Init("aleks34", "aleks", "4567", "4567")
	w := testutils.SendForm(d.T(), d.Db, "POST", "/signup", r)
	d.Assertions.Equal(201, w.Result().StatusCode)
}

func (d *DeleteAnswerSuite) AddQuestion() {
	var q models.Question
	date := time.Now().Format("02.01.2006")
	q.Init(date, "How")
	w := testutils.SendForm(d.T(), d.Db, "POST", "/questions", q)
	d.Assertions.Equal(201, w.Result().StatusCode)
}

func (d *DeleteAnswerSuite) CreateAnswer() {
	var u models.User_question
	date := time.Now().Format("02.01.2006")
	time_now := time.Now().Format("02.01.2006 15:04:05")
	u.Init(date, "aleks34", "Good", time_now, time_now)
	w := testutils.SendForm(d.T(), d.Db, "POST", "/new", u)
	d.Assertions.Equal(201, w.Result().StatusCode)
}

func (d *DeleteAnswerSuite) DeleteAnswer() {
	w := testutils.SendForm(d.T(), d.Db, "DELETE", "/new", nil)
	d.Assertions.Equal(200, w.Result().StatusCode)
}

func (d *DeleteAnswerSuite) Test() {
	d.SetupSuiteDeleteAnswer()
	d.SignUp()
	d.AddQuestion()
	d.CreateAnswer()
	d.DeleteAnswer()
	d.ClearDatabaseDeleteAnswers()
}

func (d *DeleteAnswerSuite) ClearDatabaseDeleteAnswers() {
	testutils.ClearDatabase(d.Db)
}

func Test_Delete_Answer(t *testing.T) {
	suite.Run(t, new(DeleteAnswerSuite))
}
