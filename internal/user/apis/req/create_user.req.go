package usrreq

type CreateUserReq struct {
	Email    string `json:"email" validate:"email,required"` // require duoi dang email
	FullName string `json:"fullname" validate:"min=10,max=30,required"`
	Password string `json:"password" validate:"password,required"` // password is not existed in required library
}

//
