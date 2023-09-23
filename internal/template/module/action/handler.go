package action

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/response"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/validatorext"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/action/dto"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/adapter"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/httpext"
)

// Handler handles incoming requests
type Handler struct {
	service  *Service
	validate *validator.Validate
}

// NewHandler initializes a new Handler
func NewHandler(s *Service, v *validator.Validate) *Handler {
	return &Handler{service: s, validate: v}
}

// Create handles entity create post request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var v dto.CreateActionDTO
	// validate the request body
	validationErrs, err := validatorext.ParseValidateRequestBody(r.Body, &v, h.validate)
	if validationErrs != nil {
		response.RespondError(http.StatusBadRequest, constant.Errors, validationErrs, w)
		return
	}
	if err != nil {
		response.RespondError(http.StatusBadRequest, constant.Errors, []string{err.Error()}, w)
		return
	}
	d, httpErr := h.service.create(v, r.Context())
	if httpErr.Err != nil {
		response.RespondError(httpErr.Code, constant.Error, httpErr.Err.Error(), w)
		return
	}
	response.Respond(http.StatusCreated, response.BuildData(d), w)
}

func (h *Handler) ReadMany(w http.ResponseWriter, r *http.Request) {
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
	d, httpErr := h.service.readMany(limit, page, nil)
	if httpErr.Err != nil {
		response.RespondError(httpErr.Code, constant.Errors, []error{httpErr.Err}, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), w)
}

func (h *Handler) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := httpext.GetURLParam(r, constant.KeyId)
	d, httpErr := h.service.ReadOne(id, nil)
	if httpErr.Err != nil {
		response.RespondError(httpErr.Code, constant.Errors, []string{httpErr.Err.Error()}, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), w)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := httpext.GetURLParam(r, constant.KeyId)
	var v dto.UpdateActionDTO
	validationErrs, err := validatorext.ParseValidateRequestBody(r.Body, &v, h.validate)
	if validationErrs != nil {
		response.RespondError(http.StatusBadRequest, constant.Errors, validationErrs, w)
		return
	}
	if err != nil {
		response.RespondError(http.StatusBadRequest, constant.Errors, []string{err.Error()}, w)
		return
	}
	d, httpErr := h.service.Update(id, &v, nil)
	if httpErr.Err != nil {
		response.RespondError(httpErr.Code, constant.Errors, []string{httpErr.Err.Error()}, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), w)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := httpext.GetURLParam(r, constant.KeyId)
	d, httpErr := h.service.Delete(id, r.Context())
	if httpErr.Err != nil {
		response.RespondError(httpErr.Code, constant.Errors, []string{httpErr.Err.Error()}, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), w)
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	key := httpext.GetQueryParam(r, "key")
	d, httpErr := h.service.SearchByKey(key, r.Context())
	if httpErr.Err != nil {
		response.RespondError(httpErr.Code, constant.Errors, []string{httpErr.Err.Error()}, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), w)
}
