package email

type Email interface {
	// Send The body should be a valid HTML string
	Send(to []string, subject, body string) error
}
