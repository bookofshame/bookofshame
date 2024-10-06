package auth

import (
	"fmt"

	"github.com/bookofshame/bookofshame/internal/user"
	"github.com/bookofshame/bookofshame/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	cfg     config.Config
	usrRepo user.Repository
}

func NewService(cfg config.Config, r user.Repository) Service {
	return Service{
		cfg:     cfg,
		usrRepo: r,
	}
}

func (s *Service) Login(login UserLogin) (*user.User, error) {
	usr, _ := s.usrRepo.GetByPhone(login.Phone)
	if usr == nil {
		return nil, fmt.Errorf("invalid usename or password")
	}

	isValidPwd := isValidPassword(login.Password, usr.Password)
	if !isValidPwd {
		return nil, fmt.Errorf("invalid usename or password")
	}

	if !usr.IsActive {
		return nil, fmt.Errorf("inactive user")
	}

	return usr, nil
}

// Private methods
func isValidPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
