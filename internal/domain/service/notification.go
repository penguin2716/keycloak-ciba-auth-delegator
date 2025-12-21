package service

import (
	"context"
	"delegator/internal/domain/model"
)

type NotificationService interface {
	DelegationApproved(ctx context.Context, m *model.Delegation) error
	DelegationCancelled(ctx context.Context, m *model.Delegation) error
	DelegationUnauthorized(ctx context.Context, m *model.Delegation) error
}
