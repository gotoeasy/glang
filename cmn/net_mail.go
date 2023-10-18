package cmn

import (
	"gopkg.in/gomail.v2"
)

// 邮件发送配置
type MailOptions struct {
	From         string // 发件人邮箱，默认等同邮件SMTP服务器用户名
	SmtpHost     string // 邮件SMTP服务器地址，例如 smtp.163.com
	SmtpPort     int    // 邮件SMTP服务器地址端口，默认25
	SmtpUserName string // 邮件SMTP服务器用户名
	SmtpPassword string // 邮件SMTP服务器密码
}

// 邮件发送器
type MailSender struct {
	opt *MailOptions
}

// 创建邮件发送器
func NewMailSender(opt *MailOptions) *MailSender {
	if opt == nil {
		return nil
	}
	o := &MailOptions{
		From:         opt.From,
		SmtpHost:     opt.SmtpHost,
		SmtpPort:     opt.SmtpPort,
		SmtpUserName: opt.SmtpUserName,
		SmtpPassword: opt.SmtpPassword,
	}
	if o.From == "" {
		o.From = opt.SmtpUserName
	}
	if o.SmtpPort <= 0 {
		o.SmtpPort = 25
	}
	return &MailSender{opt: o}
}

// 发送邮件
func (m *MailSender) SendMail(to string, subject string, bodyHtml string, attachs []string) error {
	// 邮件服务器的配置
	mailer := gomail.NewMessage()
	mailer.SetHeader("To", to)
	mailer.SetHeader("From", m.opt.From)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", bodyHtml) // HTML内容

	// 附件
	if attachs != nil {
		for i := 0; i < len(attachs); i++ {
			attach := attachs[i]
			if IsExistFile(attach) {
				mailer.Attach(attach)
			}
		}
	}

	// 发送邮件
	dialer := gomail.NewDialer(m.opt.SmtpHost, m.opt.SmtpPort, m.opt.SmtpUserName, m.opt.SmtpPassword)
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}
	return nil
}
