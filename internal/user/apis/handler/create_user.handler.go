// handler la noi nhan va dua data den business (usecase)
package usrhandler

import (
	"context"
	usrmapper "go-ecommerce/internal/user/apis/mapper"
	usrreq "go-ecommerce/internal/user/apis/req"
	usrentity "go-ecommerce/internal/user/business/entity"
	usrusecase "go-ecommerce/internal/user/business/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateUserController struct {
	ValidatorRequest  *validator.Validate
	CreateUserUsecase usrusecase.UserCreator // need to validate
}

func (c CreateUserController) processCreateUser(
	ctx context.Context,
	req usrreq.CreateUserReq,
) (*usrentity.User, error) {
	return c.CreateUserUsecase.Execute(ctx, usrmapper.TransformCreateReqToEntity(req))
}

func (h UserHandler) HandlecreateUser(c *gin.Context) {
	var createUserReq usrreq.CreateUserReq

	if err := c.ShouldBind(&createUserReq); err != nil {
		panic(err)
	}

	if err := h.CreateUserController.ValidatorRequest.Struct(createUserReq); err != nil {
		panic(err)
	}

	user, err := h.CreateUserController.processCreateUser(
		c.Request.Context(),
		createUserReq,
	)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
