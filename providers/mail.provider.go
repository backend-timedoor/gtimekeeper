package providers

import (
	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/mail"
)

type MailServiceProvider struct{}

func (p *MailServiceProvider) Boot() {}

func (p *MailServiceProvider) Register() {
	app.Mail = mail.New(&mail.Config{
		Host:         app.Config.GetString("MAIL_HOST", "smtp.mailtrap.io"),
		Port:         app.Config.GetInt("MAIL_PORT", 2525),
		Username:     app.Config.GetString("MAIL_USERNAME", ""),
		Password:     app.Config.GetString("MAIL_PASSWORD", ""),
		RootPath:     app.Config.GetString("path.root"),
		TemplatePath: app.Config.GetString("path.mail"),
	})
}
