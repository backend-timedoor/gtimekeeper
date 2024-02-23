package email

import (
	"github.com/backend-timedoor/gtimekeeper-framework/base/mail"
	"github.com/jordan-wright/email"
)

type ExampleMail struct {
	WithQueue   bool
	SendTo      mail.SendTo
	Attachments []*email.Attachment
}

func (m *ExampleMail) From() string {
	return "Edwin Diradinata <edwindiradinata@gmail.com>"
}

func (m *ExampleMail) View() string {
	return "example.html"
}

func (m *ExampleMail) Content(data any) mail.Content {
	return mail.Content{
		Subject: "Awesome Subject {Data Here}",
		ReplyTo: []string{"edwindiradinata@gmail.com"},
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
	}
}
