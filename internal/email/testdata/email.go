package testdata

import (
	"github.com/xebia/go-exercise/internal/email"
)

var _ email.Service = (*Mock)(nil)

type Mock struct {
	SendEmailFunc func(emailAddress, subject, body string) error
}

func (m *Mock) SendEmail(emailAddress, subject, body string) error {
	return m.SendEmailFunc(emailAddress, subject, body)
}
