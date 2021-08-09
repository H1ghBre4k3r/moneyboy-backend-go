package mail

import (
	"fmt"
	"net/smtp"
	"reflect"
)

type mailMetaInfo struct {
	auth     smtp.Auth
	hostname string
}

type Mail struct {
	from    string
	to      string
	subject string
	content string
	meta    *mailMetaInfo
}

// Create a new mail object, which can be filled with relevant information
func New(username string, password string, host string, port string) *Mail {
	mail := &Mail{
		meta: &mailMetaInfo{
			auth:     smtp.PlainAuth("", username, password, host),
			hostname: fmt.Sprintf("%v:%v", host, port),
		},
	}
	return mail
}

// Specify the sender of this mail
func (m *Mail) From(sender string) *Mail {
	m.from = sender
	return m
}

// Specify the recipient of this mail
func (m *Mail) To(recipient string) *Mail {
	m.to = recipient
	return m
}

// Specify the subject of this mail
func (m *Mail) Subject(subject string) *Mail {
	m.subject = subject
	return m
}

// Specify the content of this mail
func (m *Mail) Content(content string) *Mail {
	m.content = content
	return m
}

// Send this mail.
// returns an errror, if not all required fields are filled or the sending was unsuccessful
func (m *Mail) Send() error {
	if err := m.check(); err != nil {
		return err
	}
	msg := []byte(fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\n%s", m.to, m.from, m.subject, m.content))
	recipients := []string{m.to}
	err := smtp.SendMail(m.meta.hostname, m.meta.auth, m.from, recipients, msg)
	return err
}

// Check, if all fields in this struct are filled with valid values.
func (m *Mail) check() error {
	var err error
	val := reflect.Indirect(reflect.ValueOf(m))
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).IsZero() {
			err = fmt.Errorf("field '%s' may not be empty", val.Type().Field(i).Name)
			break
		}
	}
	return err
}
