package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user/dto"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/adapter"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/httppkg"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/jsonpkg"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/response"
	validatorpkg "github.com/tanveerprottoy/stdlib-go-template/pkg/validator"
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

func (h *Handler) parseValidateRequestBody(r *http.Request) (dto.CreateUpdateUserDTO, error) {
	var d dto.CreateUpdateUserDTO
	defer r.Body.Close()
	err := jsonpkg.Decode(r.Body, &d)
	if err != nil {
		return d, err
	}
	// validate request body
	err = h.validate.Struct(d)
	if err != nil {
		// Struct is invalid
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	}
	return d, err
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// d, err := h.parseValidateRequestBody(r)
	var d dto.CreateUpdateUserDTO
	validationErrs, err := validatorpkg.ParseValidateRequestBody(r.Body, &d, h.validate)
	if validationErrs != nil {
		response.RespondError(http.StatusBadRequest, constant.Errors, validationErrs, w)
		return
	}
	if err != nil {
		response.RespondError(http.StatusBadRequest, constant.Error, err.Error(), w)
		return
	}
	ctx := r.Context()
	e, httpErr := h.service.Create(&d, ctx)
	if httpErr != nil {
		response.RespondError(httpErr.Code, constant.Error, httpErr.Err.Error(), w)
		return
	}
	response.Respond(http.StatusCreated, e, w)
}

func (h *Handler) ReadMany(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 1
	var err error
	limitStr := httppkg.GetQueryParam(r, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, constant.Error, err.Error(), w)
			return
		}
	}
	pageStr := httppkg.GetQueryParam(r, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, constant.Error, err.Error(), w)
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
	id := httppkg.GetURLParam(r, constant.KeyId)
	e, httpErr := h.service.ReadOne(id, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, constant.Error, httpErr.Err.Error(), w)
		return
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	d, err := h.parseValidateRequestBody(r)
	if err != nil {
		response.RespondError(http.StatusBadRequest, constant.Errors, err, w)
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
	id := httppkg.GetURLParam(r, constant.KeyId)
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
