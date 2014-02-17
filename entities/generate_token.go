package entities

import jwt "github.com/dgrijalva/jwt-go"

func GenerateToken(key []byte, uid string) string {
    token := jwt.New(jwt.GetSigningMethod("HS256"))
    token.Claims["UID"] = uid
    tokenString, _ := token.SignedString(key)
    return tokenString
}
