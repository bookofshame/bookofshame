package offender

import (
	"fmt"
	"io"

	"github.com/bookofshame/bookofshame/pkg/config"
	"github.com/bookofshame/bookofshame/pkg/constants"
	"github.com/bookofshame/bookofshame/pkg/storage"
)

type Service struct {
	cfg          config.Config
	offenderRepo Repository
	storage      storage.Storage
}

func NewService(cfg config.Config, offenderRepo Repository, storage storage.Storage) Service {
	return Service{
		cfg:          cfg,
		offenderRepo: offenderRepo,
		storage:      storage,
	}
}

func (s *Service) GetAll() ([]Offender, error) {
	offenders, err := s.offenderRepo.GetAll()

	if err != nil {
		return []Offender{}, err
	}

	return offenders, nil
}

func (s *Service) Create(offender Offender, photo io.Reader) error {
	if offender.FullName == "" || offender.Address == "" {
		return fmt.Errorf("name and address is required")
	}

	exists, err := s.offenderRepo.AlreadyExists(offender.FullName, offender.DistrictId)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("offender already exists")
	}

	if photo != nil {
		photoName := offender.GeneratePhotoName()
		filePath := getPhotoFilePath(photoName)
		offender.Photo = &photoName

		err = s.storage.Upload(photo, filePath)
		if err != nil {
			return err
		}
	}

	_, err = s.offenderRepo.Create(offender)

	if err != nil {
		err := s.storage.Delete([]string{getPhotoFilePath(*offender.Photo)})
		if err != nil {
			return err
		}

		return err
	}

	return nil
}

func (s *Service) Delete(id int) error {
	offender, err := s.offenderRepo.Get(id)
	if err != nil {
		return err
	}

	err = s.offenderRepo.Delete(id)
	if err != nil {
		return err
	}

	if offender.Photo != nil {
		filePath := getPhotoFilePath(*offender.Photo)
		err := s.storage.Delete([]string{filePath})
		if err != nil {
			return err
		}
	}

	return nil
}

func getPhotoFilePath(image string) string {
	return constants.SotrageDirName + "/" + image
}
