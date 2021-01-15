package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

var (
	expirationDuration = 24 * time.Hour

	ErrTokenInvalid       = errors.New("token is invalid")
	ErrTokenInvalidIssuer = errors.New("token is from different issuer")
	ErrTokenExpired       = errors.New("token is expired")
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
		return "", fmt.Errorf("jwt#SignedString error: %s", err)
	}
	return signedToken, nil
}

func (i *TokenIssuer) Verify(token string) (*Claims, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return i.signedKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			}
		}
		log.Debugf("jwt#ParseWithClaims error %s", err)
		return nil, ErrTokenInvalid
	}
	if claims.Valid() != nil {
		log.Debugf("jwt#ParseWithClaims error %s", err)
		return nil, ErrTokenInvalid
	}
	if claims.StandardClaims.Issuer != i.name {
		return nil, ErrTokenInvalidIssuer
	}
	return &claims, nil
}
