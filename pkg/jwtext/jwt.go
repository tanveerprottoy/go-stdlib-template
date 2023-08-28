package jwtext

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/config"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/timeext"
)

// Create the JWT key used to create the signature
var JwtKey = []byte(config.GetEnvValue("secret"))

// GenerateToken generates a new token
func GenerateToken(payload map[string]any) string {
	/* RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(datetime.Now()),
		NotBefore: jwt.NewNumericDate(datetime.Now()),
		Issuer:    "test",
		Subject:   "somebody",
		ID:        "1",
		Audience:  []string{"somebody_else"},
	}, */
	// Declare the expiration time of the token
	// token := jwt.New(jwt.SigningMethodRS256)
	claims := jwt.MapClaims{}
	claims["exp"] = jwt.NewNumericDate(timeext.AddDate(0, 0, 3))
	claims["authorized"] = true
	claims["id"] = payload["id"]
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JwtKey)
	return tokenString
}

func Parse(tokenBody string, claims jwt.MapClaims) (jwt.MapClaims, error) {
	token, err := jwt.Parse(
		tokenBody,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(""), nil
		},
	)
	if err != nil {
		return nil, errors.New("malformed token")
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
