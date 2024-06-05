package repository

import (
	"context"

	"echo-model/internal/domain/model/aggregates"
)

type User interface {
	GetProfile(ctx context.Context, id string) (*aggregates.UserInfo, error)
}
