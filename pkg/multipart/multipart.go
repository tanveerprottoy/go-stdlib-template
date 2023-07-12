package multipart

import (
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/file"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/httppkg"
)

func ParseMultipartForm(maxMemory int64, r *http.Request) error {
	return r.ParseMultipartForm(maxMemory)
}

func RetrieveSaveFile(key, rootDir, destFileName string, r *http.Request) (string, error) {
	// Retrieve the file from form data
	f, h, err := httppkg.GetFile(r, key)
	if err != nil {
		return "", err
	}
	defer f.Close()
	p, err := file.SaveFile(f, rootDir, destFileName+filepath.Ext(h.Filename))
	if err != nil {
		return "", err
	}
	return p, nil
}

func HandleFileForKey(key string, rootDir, destFileName string, r *http.Request) (string, error) {
	return RetrieveSaveFile(key, rootDir, destFileName, r)
}

func HandleFilesForKeys(keys []string, rootDir, destFileName string, r *http.Request) ([]string, error) {
	var paths []string
	for _, k := range keys {
		p, err := RetrieveSaveFile(k, rootDir, destFileName, r)
		if err != nil {
			return paths, err
		}
		paths = append(paths, p)
	}
	return paths, nil
}

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
