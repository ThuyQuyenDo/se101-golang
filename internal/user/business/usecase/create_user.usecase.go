// xu ly va check data
package usrusecase

import (
	"context"
	"errors"
	usrentity "go-ecommerce/internal/user/business/entity"
	utils "go-ecommerce/utils/salt"
)

type UserCreator interface {
	Execute(
		ctx context.Context,
		entity usrentity.User,
	) (*usrentity.User, error)
}

// func need to be used
type CreateUserRepo interface {
	CreateUser(ctx context.Context, entity usrentity.User) (*usrentity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*usrentity.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type createUserUsecase struct {
	repo   CreateUserRepo
	hasher Hasher
}

func (biz *createUserUsecase) Execute(
	ctx context.Context,
	entity usrentity.User,
) (*usrentity.User, error) {
	oldUser, _ := biz.repo.FindUserByEmail(ctx, entity.Email)

	if oldUser != nil {
		return nil, errors.New("email is existed")
	}

	/*
		abcd1234 -> hash = 3707185661346832114kjdbfd
		abcd1234 -> hash = 3707185661346832114kjdbfd => nguy hiem

		abcd1234 -> hash + "random string" = 3707185661346832114kjdbfd
		-> hash(pass + salt) ==
	*/

	salt := utils.GenSalt(20)
	entity.Password = biz.hasher.Hash(entity.Password + salt)
	entity.Salt = salt

	user, err := biz.repo.CreateUser(ctx, entity)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserCreator(
	repo CreateUserRepo,
	hash Hasher,
) UserCreator {
	return &createUserUsecase{
		repo:   repo,
		hasher: hash,
	}
}
