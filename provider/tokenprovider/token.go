package tokenprovider

import (
	"time"
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserID *string `json:"user_id"`
	Role   string  `json:"role"`
	Type   string  `json:"type"`
}
