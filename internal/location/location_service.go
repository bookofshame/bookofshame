package location

import "github.com/bookofshame/bookofshame/pkg/config"

type Service struct {
	cfg          config.Config
	locationRepo Repository
}

func NewService(cfg config.Config, locationRepo Repository) Service {
	return Service{
		cfg:          cfg,
		locationRepo: locationRepo,
	}
}

func (s Service) GetDivisions() ([]Division, error) {
	divisions, err := s.locationRepo.GetDivisions()
	if err != nil {
		return nil, err
	}

	return divisions, nil
}

func (s Service) GetDistricts(divisionId int) ([]District, error) {
	districts, err := s.locationRepo.GetDistricts(divisionId)
	if err != nil {
		return nil, err
	}

	return districts, nil
}

func (s Service) GetUpazilas(districtId int) ([]Upazila, error) {
	upazilas, err := s.locationRepo.GetUpazilas(districtId)
	if err != nil {
		return nil, err
	}

	return upazilas, nil
}

func (s Service) GetUnions(upazilaId int) ([]Union, error) {
	unions, err := s.locationRepo.GetUnions(upazilaId)
	if err != nil {
		return nil, err
	}

	return unions, nil
}
