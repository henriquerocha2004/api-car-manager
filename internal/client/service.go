package client

import (
	"context"
	"time"

	"github.com/henriquerocha2004/api-car-manager/pkg/document"
	"github.com/oklog/ulid/v2"
)

type ClientService interface {
	Create(ctx context.Context, dto *ClientRequestDto) error
	Update(ctx context.Context, id ulid.ULID, dto *ClientRequestDto) error
	Delete(ctx context.Context, id ulid.ULID) error
	FindOne(ctx context.Context, id ulid.ULID) (*Client, error)
	FindByCriteria(ctx context.Context, criteria any) ([]Client, error)
}

type Service struct {
	repository Repository
}

func NewClientService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, dto *ClientRequestDto) error {

	birth, _ := time.Parse("2006-")

	client := Client{
		Name:       dto.Name,
		LastName:   dto.LastName,
		EntityType: entityType(dto.EntityType),
		Document:   document.Document(dto.Document),
	}
}

func (s *Service) Update(ctx context.Context, id ulid.ULID, dto *ClientRequestDto) error {
	return nil
}

func (s *Service) Delete(ctx context.Context, id ulid.ULID) error {
	return nil
}
