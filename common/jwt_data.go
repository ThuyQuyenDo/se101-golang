package common

const (
	CurrentRequester = ""
)

type Requester interface {
	GetUserId() string
	GetRole() string
}

var _ Requester = (*JWTUserData)(nil)

// JWTUserData
type JWTUserData struct {
	ID   string
	Role string `json:"role" gorm:"column:role;"`
}

func (u JWTUserData) GetUserId() string {
	return u.ID
}

func (u JWTUserData) GetRole() string {
	return u.Role
}
