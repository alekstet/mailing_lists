package process_test

import (
	"testing"
	"time"

	"github.com/alekstet/mailing_lists/sender/api/models"
	"github.com/alekstet/mailing_lists/sender/testutils"

	"github.com/stretchr/testify/suite"
)

type AddQuestionSuite struct {
	testutils.Store
}

func (a *AddQuestionSuite) SetupSuiteAddQuestion() {
	a.Db = testutils.LoadDatabase()
}

func (a *AddQuestionSuite) AddQuestion() {
	var r models.Question
	date := time.Now().Format("02.01.2006")
	r.Init(date, "How")
	w := testutils.SendForm(a.T(), a.Db, "POST", "/questions", r)
	a.Assertions.Equal(201, w.Result().StatusCode)
}

func (a *AddQuestionSuite) ExistsAddQuestion() {
	var r models.Question
	date := time.Now().Format("02.01.2006")
	r.Init(date, "How")
	w := testutils.SendForm(a.T(), a.Db, "POST", "/questions", r)
	a.Assertions.Equal(400, w.Result().StatusCode)
}

func (a *AddQuestionSuite) Test() {
	a.SetupSuiteAddQuestion()
	a.AddQuestion()
	a.ExistsAddQuestion()
	a.ClearDatabaseAddQuestion()
}

func (a *AddQuestionSuite) ClearDatabaseAddQuestion() {
	testutils.ClearDatabase(a.Db)
}

func Test_Add_Question(t *testing.T) {
	suite.Run(t, new(AddQuestionSuite))
}
