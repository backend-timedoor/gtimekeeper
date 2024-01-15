package mail

import (
	envelop "github.com/backend-timedoor/gtimekeeper-framework/utils/app/email"
	"github.com/jordan-wright/email"
)

type ExampleMail struct {
	SendTo      envelop.SendTo
	Attachments []*email.Attachment
}

func (m *ExampleMail) From() string {
	return "Edwin Diradinata <edwindiradinata@gmail.com>"
}

func (m *ExampleMail) View() string {
	return "example.html"
}

func (m *ExampleMail) Content(data any) envelop.Content {
	return envelop.Content{
		Subject: "Awesome Subject {Data Here}",
		ReplyTo: []string{"edwindiradinata@gmail.com"},
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
	}
}
