package usecase

import (
	"context"
	"delegator/internal/domain/model"
	"delegator/internal/domain/repository"
)

type SettingsUsecase struct {
	repo repository.SettingsRepository
}

func NewSettingsUsecase(repo repository.SettingsRepository) *SettingsUsecase {
	return &SettingsUsecase{
		repo: repo,
	}
}

func (u *SettingsUsecase) Load(ctx context.Context) (*model.Settings, error) {
	return u.repo.Load(ctx)
}

func (u *SettingsUsecase) Save(ctx context.Context, m *model.Settings) (*model.Settings, error) {
	return u.repo.Save(ctx, m)
}
