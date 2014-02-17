package entities

import "code.google.com/p/go.crypto/bcrypt"

func CompareSecret(encrypted_secret, plain_secret []byte) bool {
    if err := bcrypt.CompareHashAndPassword(encrypted_secret, plain_secret); err != nil {
        return false
    }

    return true
}
