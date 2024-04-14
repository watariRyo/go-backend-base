package repository

import (
	"context"

	"github.com/watariRyo/baby-map-server/domain/model"
)

type UsersRepository interface {
	GetUserByUUID(ctx context.Context, conn DBConnection, uuid string) (*model.User, error)
	Insert(ctx context.Context, conn DBConnection, input *model.User) (*model.User, error)
	Exist(ctx context.Context, conn DBConnection, uuid string) bool
}
