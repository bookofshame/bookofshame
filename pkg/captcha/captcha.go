package captcha

type Captcha interface {
	Verify(token string) (bool, error)
}
