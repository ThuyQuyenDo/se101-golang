package usrentity

import "go-ecommerce/common"

type User struct {
	common.BaseEntity `json:",inline"`
	Email             string        `json:"email" gorm:"column:email;"`
	FullName          string        `json:"fullname" gorm:"column:fullname;"`
	Phone             string        `json:"phone" gorm:"column:phone;"`
	Role              string        `json:"role" gorm:"column:role;"`
	Addr              string        `json:"addr" gorm:"column:addr;"`
	Password          string        `json:"-" gorm:"column:password;"`
	Salt              string        `json:"-" gorm:"column:salt;"`
	Avatar            *common.Image `json:"avatar" gorm:"column:avatar;"`
}

//Hook
func (User) TableName() string {
	return "users"
}

//entity quan trong nhat
