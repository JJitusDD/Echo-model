package repository

import (
	"context"

	"echo-model/internal/domain/model/aggregates"
	"echo-model/internal/domain/model/entities"
	"echo-model/internal/domain/model/request"
	"echo-model/internal/domain/model/response"
)

type User interface {
	GetProfile(ctx context.Context, id string) (*aggregates.UserInfo, error)
	Create(ctx context.Context, req entities.User) (res *entities.User, err error)
	Update(ctx context.Context, billID string, req entities.User) (res *entities.User, err error)
	UpdateWithSelect(ctx context.Context, ID string, selector []string, req entities.User) (res *entities.User, err error)
	UpdateWithCondition(ctx context.Context, req entities.User, keysAndValues ...string) (err error)
	FindAll(ctx context.Context, in map[string]interface{}, paging *request.Pagination) (res []*entities.User, pageRes *response.PaginationResponse, err error)
	FindById(ctx context.Context, id string) (res *entities.User, err error)
	Find(ctx context.Context, keysAndValues ...string) (ou []*entities.User, err error)
	FindOne(ctx context.Context, keysAndValues ...string) (ou *entities.User, err error)
	FindCustomerID(ctx context.Context, in map[string]interface{}) string
	Delete(ctx context.Context, id string) (e error)
}
