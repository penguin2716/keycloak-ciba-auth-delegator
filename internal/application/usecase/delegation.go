package usecase

import (
	"context"
	"delegator/internal/domain/model"
	"delegator/internal/domain/repository"
	"delegator/internal/domain/service"

	"github.com/rs/zerolog/log"
)

type DelegationUsecase struct {
	repo     repository.DelegationRepository
	notifier service.NotificationService
}

func NewDelegationUsecase(repo repository.DelegationRepository, ns service.NotificationService) *DelegationUsecase {
	return &DelegationUsecase{
		repo:     repo,
		notifier: ns,
	}
}

// Delegation作成時に指定するパラメータ
type CreateDelegationInput struct {
	AcrValues       string
	BindingMessage  string
	ConsentRequired bool
	LoginHint       string
	Scope           string
	AuthToken       string
}

func (u *DelegationUsecase) List(ctx context.Context, limit int, offset int) ([]*model.Delegation, error) {
	// limitは負の値にならない
	if limit < 0 {
		limit = 20
	}
	// 同時に取得できる最大レコード数は100
	if limit > 100 {
		limit = 100
	}
	// offsetは負の値にならない
	if offset < 0 {
		offset = 0
	}
	// リポジトリから一括取得して返却
	return u.repo.List(ctx, limit, offset)
}

func (u *DelegationUsecase) Create(ctx context.Context, input *CreateDelegationInput) (*model.Delegation, error) {
	// リクエストパラメータからモデルを作成
	m, err := model.NewDelegation(&model.NewDelegationArgs{
		AcrValues:       input.AcrValues,
		BindingMessage:  input.BindingMessage,
		ConsentRequired: input.ConsentRequired,
		LoginHint:       input.LoginHint,
		Scope:           input.Scope,
		AuthToken:       input.AuthToken,
	})
	if err != nil {
		return nil, err
	}
	// 作成した結果を返却
	return u.repo.Create(ctx, m)
}

func (u *DelegationUsecase) GetById(ctx context.Context, id string) (*model.Delegation, error) {
	// 指定したIDのモデルを取得して返却
	return u.repo.GetById(ctx, id)
}

func (u *DelegationUsecase) ApproveById(ctx context.Context, id string) (*model.Delegation, error) {
	// 指定したIDのモデルを取得
	m, err := u.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	// Approveする
	if err := m.Approve(); err != nil {
		return nil, err
	}

	// Approveされたことを通知する（非同期）
	go func() {
		err := u.notifier.DelegationApproved(context.Background(), m)
		if err != nil {
			log.Warn().Err(err).Send()
		}
	}()

	// Approveした結果を保存
	return u.repo.UpdateById(ctx, id, m)
}

func (u *DelegationUsecase) CancelById(ctx context.Context, id string) (*model.Delegation, error) {
	// 指定したIDのモデルを取得
	m, err := u.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	// Cancelする
	if err := m.Cancel(); err != nil {
		return nil, err
	}

	// Cancelされたことを通知する（非同期）
	go func() {
		err := u.notifier.DelegationCancelled(context.Background(), m)
		if err != nil {
			log.Warn().Err(err).Send()
		}
	}()

	// Cancelした結果を保存
	return u.repo.UpdateById(ctx, id, m)
}

func (u *DelegationUsecase) UnauthorizeById(ctx context.Context, id string) (*model.Delegation, error) {
	// 指定したIDのモデルを取得
	m, err := u.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	// Unauthorizeする
	if err := m.Unauthorize(); err != nil {
		return nil, err
	}

	// Unauthorizeされたことを通知する（非同期）
	go func() {
		err := u.notifier.DelegationUnauthorized(context.Background(), m)
		if err != nil {
			log.Warn().Err(err).Send()
		}
	}()

	// Unauthorizeした結果を保存
	return u.repo.UpdateById(ctx, id, m)
}

func (u *DelegationUsecase) DeleteById(ctx context.Context, id string) error {
	// 指定したIDのモデルを削除
	return u.repo.DeleteById(ctx, id)
}
