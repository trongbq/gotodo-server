package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	expirationDuration = 5 * time.Minute

	ErrTokenInvalid = errors.New("Authentication token is invalid")
)

type TokenIssuer struct {
	signedKey []byte
	name      string
}

func NewTokenIssuer(key []byte, issuer string) *TokenIssuer {
	return &TokenIssuer{
		signedKey: key,
		name:      issuer,
	}
}

type Claims struct {
	UserID int64
	jwt.StandardClaims
}

func (i *TokenIssuer) Generate(userID int64) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationDuration).Unix(),
			Issuer:    i.name,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(i.signedKey)
	if err != nil {
		return "", fmt.Errorf("JWT#SignedString error: %s", err)
	}
	return signedToken, nil
}

func (i *TokenIssuer) Verify(token string) (Claims, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return i.signedKey, nil
	})
	if err != nil || claims.Valid() != nil || claims.StandardClaims.Issuer != i.name {
		return claims, ErrTokenInvalid
	}
	return claims, nil
}
