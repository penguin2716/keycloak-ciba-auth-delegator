package repository

import (
	"context"
	"delegator/internal/domain/model"
)

type DelegationRepository interface {
	List(ctx context.Context, limit int, offset int) ([]*model.Delegation, error)
	Create(ctx context.Context, m *model.Delegation) (*model.Delegation, error)
	GetById(ctx context.Context, id string) (*model.Delegation, error)
	UpdateById(ctx context.Context, id string, m *model.Delegation) (*model.Delegation, error)
	DeleteById(ctx context.Context, id string) error
}
