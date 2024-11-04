package testdata

import (
	"context"
	"github.com/xebia/go-exercise/internal/email"
)

var _ email.Service = (*Mock)(nil)

type Mock struct {
	SendEmailFunc func(ctx context.Context, emailAddress, subject, body string) error
}

func (m *Mock) SendEmail(ctx context.Context, emailAddress, subject, body string) error {
	return m.SendEmailFunc(ctx, emailAddress, subject, body)
}
