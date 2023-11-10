package usrhandler

import (
	usrusecase "go-ecommerce/internal/user/business/usecase"
	usrrepo "go-ecommerce/internal/user/repository"
	"go-ecommerce/provider/hasher"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	CreateUserController
}

func NewUserHandler(
	validatorRequest *validator.Validate,
	repo usrrepo.UserRepository,
	hasher *hasher.MD5Hash,
) *UserHandler {
	return &UserHandler{
		CreateUserController: CreateUserController{
			ValidatorRequest: validatorRequest,
			CreateUserUsecase: usrusecase.NewUserCreator(
				repo,
				hasher,
			),
		},
	}
}
