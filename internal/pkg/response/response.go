package response

import (
	"encoding/json"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
)

type Response[T any] struct {
	Data any `json:"data"`
}

func writeResponse(w http.ResponseWriter, b []byte) (int, error) {
	return w.Write(b)
}

func BuildData[T any](payload T) *Response[T] {
	return &Response[T]{Data: payload}
}

func Respond(code int, payload any, w http.ResponseWriter) {
	res, err := json.Marshal(payload)
	if err != nil {
		RespondError(http.StatusInternalServerError, "error", err, w)
		return
	}
	w.WriteHeader(code)
	writeResponse(w, res)
}

func RespondError(code int, key string, err any, w http.ResponseWriter) {
	w.WriteHeader(code)
	res, err := json.Marshal(map[string]any{key: err})
	if err != nil {
		// log failed to marshal
		writeResponse(w, []byte(constant.InternalServerError))
		return
	}
	writeResponse(w, res)
}

func RespondErrorMessage(code int, msg string, w http.ResponseWriter) {
	w.WriteHeader(code)
	res, err := json.Marshal(map[string]string{"error": msg})
	if err != nil {
		writeResponse(w, []byte(err.Error()))
		return
	}
	writeResponse(w, res)
}

func RespondAlt(code int, payload any, w http.ResponseWriter) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		RespondError(http.StatusInternalServerError, "error", err, w)
	}
}

func RespondErrorAlt(code int, errMsg string, w http.ResponseWriter) {
	http.Error(w, errMsg, code)
}
