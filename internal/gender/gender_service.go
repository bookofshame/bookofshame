package gender

import "github.com/bookofshame/bookofshame/pkg/config"

type Service struct {
	cfg        config.Config
	genderRepo Repository
}

func NewService(cfg config.Config, genderRepo Repository) Service {
	return Service{
		cfg:        cfg,
		genderRepo: genderRepo,
	}
}

func (s *Service) GetAll() ([]Gender, error) {
	genders, err := s.genderRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return genders, nil
}
