package main

import (
	"crypto/md5"
	"encoding/hex"
)

//---------------cryptage mdp---------------
func crypt(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
