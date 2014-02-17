package entities

import "code.google.com/p/go.crypto/bcrypt"

var HashStrength = 10

func EncryptSecret(secret []byte) []byte {
    hash, _ := bcrypt.GenerateFromPassword(secret, HashStrength)
    return hash
}
