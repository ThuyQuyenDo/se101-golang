package usrmapper

import (
	usrreq "go-ecommerce/internal/user/apis/req"
	usrentity "go-ecommerce/internal/user/business/entity"
)

func TransformCreateReqToEntity(req usrreq.CreateUserReq) usrentity.User {
	return usrentity.User{
		Email:    req.Email,
		FullName: req.FullName,
		Password: req.Password,
	}
}
