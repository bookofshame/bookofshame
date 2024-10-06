package email

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/bookofshame/bookofshame/pkg/config"
	"github.com/bookofshame/bookofshame/pkg/logging"
	"go.uber.org/zap"
)

type Client struct {
	cfg    config.Config
	logger *zap.SugaredLogger
}

func NewEmailClient(ctx context.Context, cfg config.Config) *Client {
	return &Client{
		cfg:    cfg,
		logger: logging.FromContext(ctx),
	}
}

func (s Client) getHeaders(to []string, subject string) string {
	fromHeader := fmt.Sprintf("From: Book Of Shame <%s>\r\n", s.cfg.SmtpFromEmail)
	toHeader := fmt.Sprintf("To: %s\r\n", strings.Join(to, ","))
	subjectHeader := fmt.Sprintf("Subject: %s\r\n", subject)
	mimeHeader := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\"\r\n"

	return strings.Join([]string{fromHeader, toHeader, subjectHeader, mimeHeader}, "")
}

func (s Client) Send(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", s.cfg.SmtpUsername, s.cfg.SmtpPassword, s.cfg.SmtpHost)

	headers := s.getHeaders(to, subject)
	message := []byte(headers + "\r\n" + body + "\r\n")

	smtpUrl := s.cfg.SmtpHost + ":" + s.cfg.SmtpPort

	err := smtp.SendMail(smtpUrl, auth, s.cfg.SmtpFromEmail, to, message)

	if err != nil {
		s.logger.Errorf("failed to send email. error: %w", err)
		return fmt.Errorf("failed to send email")
	}

	return nil
}
