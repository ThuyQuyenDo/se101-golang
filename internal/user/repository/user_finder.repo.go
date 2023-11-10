package usrrepo

import (
	"context"
	"go-ecommerce/db"
	usrentity "go-ecommerce/internal/user/business/entity"
)

type userFinderImpl struct {
	db db.Database
}

// FindUserByEmail implements UserRepoFinder.
func (u *userFinderImpl) FindUserByEmail(
	ctx context.Context, email string,
) (*usrentity.User, error) {
	userEntity := usrentity.User{}

	// Select * from " " where id = abcde
	if err := u.db.Executor.
		Where("email = ?", email).
		First(&userEntity).Error; err != nil {
		return nil, err
	}
	return &userEntity, nil
}

// FindUserByPhone implements UserRepoFinder.
func (u *userFinderImpl) FindUserByPhone(
	ctx context.Context,
	phone string,
) (*usrentity.User, error) {
	userEntity := usrentity.User{}

	// Select * from " " where id = abcde
	if err := u.db.Executor.
		Where("phone = ?", phone).
		First(&userEntity).Error; err != nil {
		return nil, err
	}
	return &userEntity, nil
}

// FindUserByID implements UserRepoFinder.
func (u *userFinderImpl) FindUserByID(
	ctx context.Context, userID string,
) (*usrentity.User, error) {
	userEntity := usrentity.User{}

	// Select * from " " where id = abcde
	if err := u.db.Executor.
		Where("id = ?", userID).
		First(&userEntity).Error; err != nil {
		return nil, err
	}
	return &userEntity, nil
}

var _ UserRepoFinder = (*userFinderImpl)(nil)

func NewUserFinderImpl(db db.Database) UserRepoFinder {
	return &userFinderImpl{
		db: db,
	}
}
