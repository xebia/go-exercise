package email

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Service interface {
	SendEmail(ctx context.Context, emailAddress, subject, body string) error
}

type service struct {
	apiKey string
}

func NewService(apiKey string) Service {
	return &service{
		apiKey: apiKey,
	}
}

func (s *service) SendEmail(ctx context.Context, recipientAddress string, subject, body string) error {
	from := mail.NewEmail("YOUR NAME", "YOUR_EMAIL@FOO.BAR")
	to := mail.NewEmail("", recipientAddress)

	htmlContent := "<strong>" + body + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, body, htmlContent)
	client := sendgrid.NewSendClient(s.apiKey)

	response, err := client.SendWithContext(ctx, message)
	if err != nil {
		return fmt.Errorf("Error sending email: %s", err)
	}
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error sending email: %d", response.StatusCode)
	}
	log.Printf("Successfully send mail to %s (%+v	)", recipientAddress, response)

	return nil
}
