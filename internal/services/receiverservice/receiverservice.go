package receiverservice

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/moura95/transferapi/config"
	"github.com/moura95/transferapi/internal/entity"
	receiverrepo "github.com/moura95/transferapi/internal/repository/receiversrepo"
	"go.uber.org/zap"
)

type Service struct {
	repository receiverrepo.IReceiverRepository
	config     config.Config
	logger     *zap.SugaredLogger
}

func NewReceiverService(repo receiverrepo.IReceiverRepository, cfg config.Config, log *zap.SugaredLogger) *Service {
	return &Service{
		repository: repo,
		config:     cfg,
		logger:     log,
	}
}

func (s *Service) Create(name, email, taxId string) error {
	dr := entity.NewReceiver(name)
	err := dr.Validate()
	if err != nil {
		return err
	}

	err = s.repository.Create(*dr)
	if err != nil {
		return fmt.Errorf("failed to create %s", err.Error())
	}
	return nil
}

func (s *Service) GetByID(uid uuid.UUID) (*entity.Receiver, error) {
	receiver, err := s.repository.GetByID(uid)

	if err != nil {
		return nil, fmt.Errorf("failed to get receiver %s", err.Error())
	}
	if receiver == nil {
		return nil, fmt.Errorf("not found")

	}

	return receiver, nil
}

func (s *Service) List() ([]entity.Receiver, error) {
	receivers, err := s.repository.GetAll()
	if err != nil {
		return []entity.Receiver{}, fmt.Errorf("failed to list receiversrepo %s", err.Error())
	}
	return receivers, nil
}

func (s *Service) Update(uid uuid.UUID, name, email, taxId string) error {
	dr := entity.NewReceiver(name)
	err := s.repository.Update(uid, dr)
	if err != nil {
		return fmt.Errorf("failed to update receiver %s", err.Error())
	}
	return nil
}
