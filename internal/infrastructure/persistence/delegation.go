package persistence

import (
	"context"
	"delegator/internal/domain/model"
	"delegator/internal/domain/repository"
	"time"

	"gorm.io/gorm"
)

// model.DelegationRepository を実装する構造体
type DelegationStore struct {
	db *gorm.DB
}

func NewDelegationStore(db *gorm.DB) *DelegationStore {
	return &DelegationStore{
		db: db,
	}
}

// DBに格納するレコードの定義
type DelegationDTO struct {
	ID              string `gorm:"primaryKey"`
	Status          string
	AcrValues       string
	BindingMessage  string
	ConsentRequired bool
	LoginHint       string
	Scope           string
	AuthToken       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// テーブル名を指定
func (o *DelegationDTO) TableName() string {
	return "delegations"
}

// DTOをドメイン層のモデルに変換する
func (o *DelegationDTO) toDomain() *model.Delegation {
	return &model.Delegation{
		ID:              o.ID,
		Status:          model.DelegationStatus(o.Status),
		AcrValues:       o.AcrValues,
		BindingMessage:  o.BindingMessage,
		ConsentRequired: o.ConsentRequired,
		LoginHint:       o.LoginHint,
		Scope:           o.Scope,
		AuthToken:       o.AuthToken,
		CreatedAt:       o.CreatedAt,
		UpdatedAt:       o.UpdatedAt,
	}
}

// ドメイン層のモデルをDTOに変換する
func toDelegationDTO(m *model.Delegation) *DelegationDTO {
	return &DelegationDTO{
		ID:              m.ID,
		Status:          string(m.Status),
		AcrValues:       m.AcrValues,
		BindingMessage:  m.BindingMessage,
		ConsentRequired: m.ConsentRequired,
		LoginHint:       m.LoginHint,
		Scope:           m.Scope,
		AuthToken:       m.AuthToken,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func (s *DelegationStore) List(ctx context.Context, limit int, offset int) ([]*model.Delegation, error) {
	// DBからすべてのレコードを取得
	records := []*DelegationDTO{}
	err := s.db.WithContext(ctx).Limit(limit).Offset(offset).Order("created_at desc").Find(&records).Error
	if err != nil {
		return nil, err
	}

	// DBから取得したレコードをドメイン層のモデルに変換して返却
	delegations := []*model.Delegation{}
	for _, record := range records {
		delegations = append(delegations, record.toDomain())
	}
	return delegations, nil
}

func (s *DelegationStore) Create(ctx context.Context, m *model.Delegation) (*model.Delegation, error) {
	// ドメイン層のモデルをDTOに変換してDBに保存
	dto := toDelegationDTO(m)
	err := s.db.WithContext(ctx).Save(dto).Error
	if err != nil {
		return nil, err
	}

	// 保存結果のDTOをドメイン層のモデルに変換して返却
	return dto.toDomain(), nil
}

func (s *DelegationStore) GetById(ctx context.Context, id string) (*model.Delegation, error) {
	// 指定されたIDのレコードを取得
	dto := &DelegationDTO{}
	err := s.db.WithContext(ctx).Where(&DelegationDTO{ID: id}).First(dto).Error
	if err == gorm.ErrRecordNotFound {
		return nil, repository.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	// 取得した結果をドメイン層のモデルに変換して返却
	return dto.toDomain(), nil
}

func (s *DelegationStore) UpdateById(ctx context.Context, id string, m *model.Delegation) (*model.Delegation, error) {
	// ドメイン層のモデルをDTOに変換して保存
	dto := toDelegationDTO(m)
	err := s.db.WithContext(ctx).Select("*").Updates(dto).Error
	if err == gorm.ErrRecordNotFound {
		return nil, repository.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	// 保存結果をドメイン層のモデルに変換して返却
	return dto.toDomain(), nil
}

func (s *DelegationStore) DeleteById(ctx context.Context, id string) error {
	// 指定されたIDのレコードを削除
	return s.db.WithContext(ctx).Delete(&DelegationDTO{ID: id}).Error
}
