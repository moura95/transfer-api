package receiverservice

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/moura95/transferapi/config"
	"github.com/moura95/transferapi/internal/entity"
	receiverrepo "github.com/moura95/transferapi/internal/repository/receiversrepo"
	"github.com/moura95/transferapi/pkg/ginx"
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

const (
	statusRascunho = "Rascunho"
	statusValidado = "Validado"
)

func (s *Service) Create(name, pixKeyType, pixKey, email, CpfCnpj string) error {
	rv := entity.NewReceiver(uuid.New(), name, pixKeyType, pixKey, email, CpfCnpj, statusRascunho)
	err := rv.Validate()
	if err != nil {
		return err
	}

	err = s.repository.Create(*rv)
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

func (s *Service) Delete(uid uuid.UUID) error {
	err := s.repository.HardDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to get receiver %s", err.Error())
	}

	return nil
}
func (s *Service) BulkDelete(uuids []string) error {
	err := s.repository.BulkDelete(uuids)
	if err != nil {
		return fmt.Errorf("failed to get receiver %s", err.Error())
	}

	return nil
}

type ListRequest struct {
	Status      string
	Name        string
	PixKeyType  string
	PixKeyValue string
	Limit       int
	Page        int
}

func (s *Service) List(filters ListRequest) ([]entity.Receiver, ginx.PageInfo, error) {
	res, err := s.repository.GetAll(
		entity.NewFilter(filters.Status, filters.Name, filters.PixKeyType, filters.PixKeyValue, filters.Limit, filters.Page))
	if err != nil {
		return []entity.Receiver{}, ginx.PageInfo{}, fmt.Errorf("failed to get receiver %s", err.Error())
	}
	pageInfo := ginx.NewPageInfo(res.Limit, res.TotalRecords, res.CurrentPage, res.TotalPages)
	return res.Receivers, pageInfo, nil
}

func (s *Service) Update(uid uuid.UUID, name, pixKeyType, pixKey, email, CpfCnpj string) error {
	rec, err := s.repository.GetByID(uid)

	if err != nil {
		return fmt.Errorf("failed to get receiver %s", err.Error())
	}
	receiver := entity.NewReceiver(uid, name, pixKeyType, pixKey, email, CpfCnpj, rec.Status)
	err = receiver.ValidateUpdate()
	if err != nil {
		return err
	}
	err = s.repository.Update(uid, receiver)
	if err != nil {
		return fmt.Errorf("failed to update receiver %s", err.Error())
	}
	return nil
}
