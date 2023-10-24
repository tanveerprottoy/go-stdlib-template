package jwk

import (
	"fmt"
	"time"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/config"
)

func Usage() error {
	jwksURL := fmt.Sprintf("%s/.well-known/jwks", config.GetEnvValue("jwkDomain"))
	jwks, err := CreateJWKS(jwksURL, 24*time.Hour, nil)
	if err != nil {
		return err
	}
	tokenBody := "fksjlksfdjlksfj"
	jwtToken, err := Parse(tokenBody, jwks.Keyfunc)
	if err != nil {
		return err
	}
	fmt.Println("jwtToken: ", *jwtToken)
	// extract the claims
	c := ParseClaims(jwtToken)
	if c == nil {
		// handle error
	}
	fmt.Println("claims: ", c)
	id := c["userId"].(string)
	fmt.Println("id: ", id)
	return nil
}
