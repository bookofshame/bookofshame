package user

import (
	"crypto/rand"
	"fmt"

	"github.com/bookofshame/bookofshame/pkg/config"
	"github.com/bookofshame/bookofshame/pkg/email"
	"github.com/bookofshame/bookofshame/pkg/sms"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	cfg         config.Config
	usrRepo     Repository
	emailClient email.Email
	smsClient   sms.Sms
}

func NewService(cfg config.Config, r Repository, emailClient email.Email, smsClient sms.Sms) Service {
	return Service{
		cfg:         cfg,
		usrRepo:     r,
		emailClient: emailClient,
		smsClient:   smsClient,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func newOtpCode() string {
	otpBytes := make([]byte, 6)
	_, err := rand.Read(otpBytes)
	if err != nil {
		return ""
	}

	for i := 0; i < 6; i++ {
		otpBytes[i] = uint8(48 + (otpBytes[i] % 10))
	}

	return string(otpBytes)
}

func (s Service) sendSmsOtp(to, otp string) error {
	message := fmt.Sprintf("Your Book of Shame OTP is %s", otp)
	err := s.smsClient.Send(sms.Payload{Number: to, Message: message})
	if err != nil {
		return fmt.Errorf("failed to send otp sms")
	}

	return nil
}

func (s Service) Get(id int) (*User, error) {
	user, err := s.usrRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s Service) GetAll() ([]User, error) {
	users, err := s.usrRepo.GetAll()

	if err != nil {
		return nil, fmt.Errorf("failed to get all user")
	}

	for i := range users {
		users[i].Normalize()
	}

	return users, nil
}

func (s Service) Create(user User) error {
	phoneExists, err := s.usrRepo.PhoneExists(user.Phone)
	if err != nil {
		return fmt.Errorf("failed to check if phone exists")
	}

	if phoneExists {
		return fmt.Errorf("phone already exists")
	}

	emailExists, err := s.usrRepo.EmailExists(user.Email)
	if err != nil {
		return fmt.Errorf("failed to check if email exists")
	}

	if emailExists {
		return fmt.Errorf("email already exists")
	}

	hashedPwd, err := hashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to create password hash")
	}

	user.Password = hashedPwd

	opt := newOtpCode()
	user.ActivationCode = &opt

	_, err = s.usrRepo.Create(user)
	if err != nil {
		return fmt.Errorf("failed to create user")
	}

	err = s.sendSmsOtp(user.Phone, *user.ActivationCode)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) VerifyPhone(code, phone string) error {
	user, err := s.usrRepo.GetByPhone(phone)
	if err != nil {
		return err
	}

	if user.ActivationCode == nil || *user.ActivationCode != code {
		return fmt.Errorf("invalid otp")
	}

	return s.usrRepo.Activate(user.Id)
}

func (s Service) VerifyEmail(code string) error {
	id, err := s.usrRepo.GetIdByActivationCode(code)

	if err != nil {
		return err
	}

	return s.usrRepo.Activate(id)
}

func (s Service) ResendOtp(phone string) error {
	user, err := s.usrRepo.GetByPhone(phone)
	if err != nil {
		return err
	}

	*user.ActivationCode = newOtpCode()
	if err := s.usrRepo.Update(*user); err != nil {
		return fmt.Errorf("failed to resend OPT")
	}

	err = s.sendSmsOtp(user.Phone, *user.ActivationCode)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) SendEmailOtp(to, otp string) error {
	body := `
    <html>
        <body>
            <p>Greetings,</p>
            <p>Your Book of Shame OTP is <b>` + otp + `.</p>
        </body>
    </html>
    `

	if err := s.emailClient.Send([]string{to}, "Book of Shame OTP", body); err != nil {
		return fmt.Errorf("failed to send otp email")
	}

	return nil
}

func (s Service) SendActivationLink(to, activationLink string) error {
	hyperlinkHtml := `<a href="` + activationLink + `">` + activationLink + `</a>`
	body := `
    <html>
        <body>
            <p>Greetings,</p>
            <p>Click on this activation link ` + hyperlinkHtml + ` to activate your account.</p>
        </body>
    </html>`

	if err := s.emailClient.Send([]string{to}, "Activate your Book of Shame", body); err != nil {
		return fmt.Errorf("failed to send activation link email")
	}

	return nil
}
