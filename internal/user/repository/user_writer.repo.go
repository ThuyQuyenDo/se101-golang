package usrrepo

import (
	"context"
	"go-ecommerce/db"
	usrentity "go-ecommerce/internal/user/business/entity"
)

type userWriteImpl struct {
	db db.Database
}

// CreateUser implements UserRepoWriter.
func (u *userWriteImpl) CreateUser(
	ctx context.Context,
	entity usrentity.User,
) (*usrentity.User, error) {
	if err := u.db.Executor.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

var _ UserRepoWriter = (*userWriteImpl)(nil)

func NewUserWriterImpl(db db.Database) UserRepoWriter {
	return &userWriteImpl{
		db: db,
	}
}
