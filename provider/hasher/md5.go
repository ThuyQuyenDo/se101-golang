package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type MD5Hash struct{}

func NewMd5Hash() *MD5Hash {
	return &MD5Hash{}
}

func (h *MD5Hash) Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data)) // ep kieu string to byte ( []byte(...) )
	return hex.EncodeToString(hasher.Sum(nil))
}
