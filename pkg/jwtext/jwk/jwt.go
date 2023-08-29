package jwk

import (
	"errors"
	"log"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Parse(tokenBody string, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenBody, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

func ParseClaims(token *jwt.Token) jwt.MapClaims {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		log.Printf("Invalid JWT Token")
		return nil
	}
}

func CreateJWKS(jwksURL string) (*keyfunc.JWKS, error) {
	// Create the keyfunc options. Use an error handler that logs. Timeout the initial JWKS refresh request after 10
	// seconds. This timeout is also used to create the initial context.Context for keyfunc.Get.
	options := keyfunc.Options{
		RefreshTimeout: time.Minute * 1,
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
	}
	// Create the JWKS from the resource at the given URL.
	return keyfunc.Get(jwksURL, options)
}
