// interact db, save final database, not check anything
package usrrepo

import (
	"context"
	usrentity "go-ecommerce/internal/user/business/entity"
)

type UserRepoFinder interface {
	FindUserByID(ctx context.Context, userID string) (*usrentity.User, error)
	FindUserByPhone(ctx context.Context, phone string) (*usrentity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*usrentity.User, error)
}

type UserRepoWriter interface {
	CreateUser(ctx context.Context, entity usrentity.User) (*usrentity.User, error)
}

type UserRepository struct {
	UserRepoFinder
	UserRepoWriter
}

func NewUserRepository(
	userFinder UserRepoFinder,
	userWriter UserRepoWriter,
) *UserRepository {
	return &UserRepository{
		UserRepoFinder: userFinder,
		UserRepoWriter: userWriter,
	}
}
