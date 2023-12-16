package model

import (
	"crypto/md5"
	"encoding/hex"
)

func (user *user) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(user.password))

	// Exchanging original user password for the encrypted version
	// to compare it with the encrypted database password
	user.password = hex.EncodeToString(hash.Sum(nil))
}