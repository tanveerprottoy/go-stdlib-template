package auth

import (
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/response"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/auth/dto"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/config"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/httpext"
)

type ServiceRemote struct {
	HTTPClient *httpext.HTTPClient
}

func NewServiceRemote(c *httpext.HTTPClient) *ServiceRemote {
	s := new(ServiceRemote)
	s.HTTPClient = c
	return s
}

func (s *ServiceRemote) Authorize(w http.ResponseWriter, r *http.Request) any {
	_, err := httpext.ParseAuthToken(r)
	if err != nil {
		response.RespondError(http.StatusForbidden, constant.Error, err, w)
		return nil
	}
	u, err := httpext.Request[dto.AuthUserDTO](
		http.MethodPost,
		fmt.Sprintf("%s%s", config.GetEnvValue("USER_SERVICE_BASE_URL"), constant.UserServiceAuthEndpoint),
		r.Header,
		nil,
		s.HTTPClient,
	)
	if err != nil {
		response.RespondError(http.StatusForbidden, constant.Error, err, w)
		return nil
	}
	return u
}
