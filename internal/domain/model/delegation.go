package model

import (
	"time"

	"github.com/google/uuid"
)

type DelegationStatus string

const (
	DelegationStatusPending      = DelegationStatus("PENDING")
	DelegationStatusSucceed      = DelegationStatus("SUCCEED")
	DelegationStatusCancelled    = DelegationStatus("CANCELLED")
	DelegationStatusUnauthorized = DelegationStatus("UNAUTHORIZED")
)

type Delegation struct {
	ID              string
	Status          DelegationStatus
	AcrValues       string
	BindingMessage  string
	ConsentRequired bool
	LoginHint       string
	Scope           string
	AuthToken       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type NewDelegationArgs struct {
	AcrValues       string
	BindingMessage  string
	ConsentRequired bool
	LoginHint       string
	Scope           string
	AuthToken       string
}

func NewDelegation(args *NewDelegationArgs) (*Delegation, error) {
	now := time.Now()
	return &Delegation{
		ID:              uuid.New().String(),
		Status:          DelegationStatusPending,
		AcrValues:       args.AcrValues,
		BindingMessage:  args.BindingMessage,
		ConsentRequired: args.ConsentRequired,
		LoginHint:       args.LoginHint,
		Scope:           args.Scope,
		AuthToken:       args.AuthToken,
		CreatedAt:       now,
		UpdatedAt:       now,
	}, nil
}

func (d *Delegation) Approve() error {
	d.Status = DelegationStatusSucceed
	return nil
}

func (d *Delegation) Cancel() error {
	d.Status = DelegationStatusCancelled
	return nil
}

func (d *Delegation) Unauthorize() error {
	d.Status = DelegationStatusUnauthorized
	return nil
}
