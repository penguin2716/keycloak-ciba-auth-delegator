package repository

import (
	"context"
	"delegator/internal/domain/model"
)

type SettingsRepository interface {
	Load(ctx context.Context) (*model.Settings, error)
	Save(ctx context.Context, tobe *model.Settings) (*model.Settings, error)
}
