package auth

type UserLogin struct {
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	CaptchaToken string `json:"captchaToken"`
}

type Token struct {
	AccessToken string `json:"accessToken"`
}
