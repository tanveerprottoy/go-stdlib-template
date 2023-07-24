package auth

import (
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/auth/dto"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/config"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/httppkg"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/response"
)

type ServiceRemote struct {
	HTTPClient *httppkg.HTTPClient
}

func NewServiceRemote(c *httppkg.HTTPClient) *ServiceRemote {
	s := new(ServiceRemote)
	s.HTTPClient = c
	return s
}

func (s *ServiceRemote) Authorize(w http.ResponseWriter, r *http.Request) any {
	_, err := httppkg.ParseAuthToken(r)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	u, err := httppkg.Request[dto.AuthUserDTO](
		http.MethodPost,
		fmt.Sprintf("%s%s", config.GetEnvValue("USER_SERVICE_BASE_URL"), constant.UserServiceAuthEndpoint),
		r.Header,
		nil,
		s.HTTPClient,
	)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	return u
}
