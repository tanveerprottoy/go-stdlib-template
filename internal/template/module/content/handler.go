package content

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/response"
	validatorpkg "github.com/tanveerprottoy/stdlib-go-template/internal/pkg/validator"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/dto"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/adapter"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/httpext"
)

// Hanlder is responsible for extracting data
// from request body and building and seding response
type Handler struct {
	service  *Service
	validate *validator.Validate
}

func NewHandler(s *Service, v *validator.Validate) *Handler {
	h := new(Handler)
	h.service = s
	h.validate = v
	return h
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var d dto.CreateUpdateContentDTO
	validationErrs, err := validatorpkg.ParseValidateRequestBody(r.Body, &d, h.validate)
	if validationErrs != nil {
		response.RespondError(http.StatusBadRequest, constant.Errors, validationErrs, w)
		return
	}
	if err != nil {
		response.RespondError(http.StatusBadRequest, constant.Error, err, w)
		return
	}
	e, httpErr := h.service.Create(&d, r.Context())
	if httpErr != nil {
		response.RespondError(httpErr.Code, constant.Error, httpErr.Err.Error(), w)
		return
	}
	response.Respond(http.StatusCreated, e, w)
}

func (h *Handler) ReadMany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("ReadMany.context.RBAC: ", ctx.Value(constant.KeyRBAC))
	fmt.Println("ReadMany.context.AuthUser: ", ctx.Value(constant.KeyAuthUser))
	limit := 10
	page := 1
	var err error
	limitStr := httpext.GetQueryParam(r, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, constant.Error, err, w)
			return
		}
	}
	pageStr := httpext.GetQueryParam(r, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, constant.Error, err, w)
			return
		}
	}
	e, httpErr := h.service.ReadMany(limit, page, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, constant.Error, httpErr.Err.Error(), w)
		return
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *Handler) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := httpext.GetURLParam(r, constant.KeyId)
	e, httpErr := h.service.ReadOne(id, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, constant.Error, httpErr.Err.Error(), w)
		return
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := httpext.GetURLParam(r, constant.KeyId)
	var d dto.CreateUpdateContentDTO
	validationErrs, err := validatorpkg.ParseValidateRequestBody(r.Body, &d, h.validate)
	if validationErrs != nil {
		response.RespondError(http.StatusBadRequest, constant.Errors, validationErrs, w)
		return
	}
	if err != nil {
		response.RespondError(http.StatusBadRequest, constant.Error, err, w)
		return
	}
	e, httpErr := h.service.Update(id, &d, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, constant.Error, httpErr.Err.Error(), w)
		return
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := httpext.GetURLParam(r, constant.KeyId)
	e, httpErr := h.service.Delete(id, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, constant.Error, httpErr.Err.Error(), w)
		return
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *Handler) Public(w http.ResponseWriter, r *http.Request) {
	response.Respond(http.StatusOK, map[string]string{"message": "public api"}, w)
}
