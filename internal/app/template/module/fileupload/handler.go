package fileupload

import (
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/core"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/multipart"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/response"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	h := new(Handler)
	h.service = service
	return h
}

func (h *Handler) parseMultipartForm(r *http.Request) error {
	return multipart.ParseMultipartForm(core.LeftShift(32, 20), r)
}

func (h *Handler) UploadOne(w http.ResponseWriter, r *http.Request) {
	err := h.parseMultipartForm(r)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	d, err := h.service.UploadOne(r)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, d, w)
}

func (h *Handler) UploadOneDisk(w http.ResponseWriter, r *http.Request) {
	err := h.parseMultipartForm(r)
	if err != nil {
		response.RespondErrorAlt(http.StatusInternalServerError, "Parse error", w)
		return
	}
	d, err := h.service.UploadOneDisk(r)
	if err != nil {
		response.RespondErrorAlt(http.StatusInternalServerError, "an error", w)
		return
	}
	response.Respond(http.StatusOK, d, w)
}

func (h *Handler) UploadMany(w http.ResponseWriter, r *http.Request) {
	/* paths, err := multipart.HandleFilesForKeys([]string{"image0, image1"}, "./uploads", "file0",r)
	if err != nil {
		response.RespondErrorAlt(http.StatusInternalServerError, "Parse error", w)
		return
	} */
	response.Respond(http.StatusOK, map[string][]string{"filePaths": {""}}, w)
}

func (h *Handler) UploadManyDisk(w http.ResponseWriter, r *http.Request) {
	err := h.parseMultipartForm(r)
	if err != nil {
		response.RespondErrorAlt(http.StatusInternalServerError, "Parse error", w)
		return
	}
	d, err := h.service.UploadManyDisk(r)
	if err != nil {
		response.RespondErrorAlt(http.StatusInternalServerError, "an error", w)
		return
	}
	response.Respond(http.StatusOK, d, w)
}

func (h *Handler) UploadManyWithKeysDisk(w http.ResponseWriter, r *http.Request) {
	err := h.parseMultipartForm(r)
	if err != nil {
		response.RespondErrorAlt(http.StatusInternalServerError, "Parse error", w)
		return
	}
	d, err := h.service.UploadManyWithKeysDisk(r)
	if err != nil {
		response.RespondErrorAlt(http.StatusInternalServerError, "an error", w)
		return
	}
	response.Respond(http.StatusOK, d, w)
}
