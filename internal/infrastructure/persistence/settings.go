package persistence

import (
	"context"
	"delegator/internal/domain/model"
	"time"

	"gorm.io/gorm"
)

type SettingsStore struct {
	db *gorm.DB
}

func NewSettingsStore(db *gorm.DB) *SettingsStore {
	return &SettingsStore{
		db: db,
	}
}

type SettingsDTO struct {
	ID              uint `gorm:"primaryKey"`
	KeycloakBaseURL string
	KeycloakRealm   string
	UpdatedAt       time.Time
}

func (o *SettingsDTO) TableName() string {
	return "settings"
}

func (o *SettingsDTO) toDomain() *model.Settings {
	return &model.Settings{
		Keycloak: model.KeycloakSettings{
			BaseURL: o.KeycloakBaseURL,
			Realm:   o.KeycloakRealm,
		},
	}
}

func toSettingsDTO(m *model.Settings) *SettingsDTO {
	return &SettingsDTO{
		ID:              1,
		KeycloakBaseURL: m.Keycloak.BaseURL,
		KeycloakRealm:   m.Keycloak.Realm,
	}
}

func (s *SettingsStore) Load(ctx context.Context) (*model.Settings, error) {
	dto := &SettingsDTO{ID: 1}
	err := s.db.WithContext(ctx).First(dto).Error
	if err != nil {
		return nil, err
	}
	return dto.toDomain(), nil
}

func (s *SettingsStore) Save(ctx context.Context, m *model.Settings) (*model.Settings, error) {
	dto := toSettingsDTO(m)
	err := s.db.WithContext(ctx).Save(dto).Error
	if err != nil {
		return nil, err
	}
	return dto.toDomain(), nil
}
