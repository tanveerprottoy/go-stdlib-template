package multipart

import (
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/file"
)

// ParseMultipartForm parses the MultipartForm
func ParseMultipartForm(maxMemory int64, r *http.Request) error {
	return r.ParseMultipartForm(maxMemory)
}

// GetFiles retrieves files from MultipartForm
func GetFiles(key string, r *http.Request) []*multipart.FileHeader {
	return r.MultipartForm.File[key]
}

// GetValues retrieves values from MultipartForm
func GetValues(r *http.Request) map[string][]string {
	return r.MultipartForm.Value
}

// GetFormFile retrieves a file from MultipartForm
func GetFormFile(key string, r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile(key)
}

// SaveFile saves the file 
func SaveFile(f multipart.File, header *multipart.FileHeader, rootDir, destFileName string, r *http.Request) (string, error) {
	defer f.Close()
	p, err := file.SaveFile(f, rootDir, destFileName+filepath.Ext(header.Filename))
	if err != nil {
		return "", err
	}
	return p, nil
}

// GetFileContentType fetches content type of the file
func GetFileContentType(file multipart.File) (string, error) {
	// to sniff the content type only the first
	// 512 bytes are used.
	buf := make([]byte, 512)

	_, err := file.Read(buf)

	if err != nil {
		return "", err
	}

	// the function that actually does the trick
	contentType := http.DetectContentType(buf)

	return contentType, nil
}
