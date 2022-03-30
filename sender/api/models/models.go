package models

type Mail struct {
	Id         int
	Text       string
	StartTime  string
	EndTime    string
	Tag        string
	MobileCode int
}

type Client struct {
	Id         int
	Phone      int
	MobileCode int
	Tag        string
}

type Message struct {
	MessageId int
	SendTime  string
	MailId    int
	ClientId  int
}
