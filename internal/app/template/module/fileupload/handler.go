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

func (h *Handler) UploadOne(w http.ResponseWriter, r *http.Request) {
	// reqMultipartParsed
	/* err := multipart.ParseMultipartForm(r)
	if err != nil {
		response.RespondErrorAlt(http.StatusInternalServerError, "Parse error", w)
		return
	}
	f, header, err := httppkg.GetFile(r, "file")
	if err != nil {
		return paths, err
	} */
}

func (h *Handler) UploadOneDisk(w http.ResponseWriter, r *http.Request) {
	err := multipart.ParseMultipartForm(core.LeftShift(32, 20), r)
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
	paths, err := multipart.HandleFilesForKeys([]string{"image0, image1"}, "./uploads", "file0",r)
	if err != nil {
		response.RespondErrorAlt(http.StatusInternalServerError, "Parse error", w)
		return
	}
	response.Respond(http.StatusOK, map[string][]string{"filePaths": paths}, w)
}

func (h *Handler) UploadManyDisk(w http.ResponseWriter, r *http.Request) {
	/* p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(p, w, r) */
}
