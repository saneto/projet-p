package security

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/twinj/uuid"
)

// TokenHash : Use for resetPassword, we hash also the token send by email
func TokenHash(text string) string {

	hasher := md5.New()
	hasher.Write([]byte(text))
	theHash := hex.EncodeToString(hasher.Sum(nil))

	//also use uuid
	u := uuid.NewV4()
	theToken := theHash + u.String()

	return theToken
}
