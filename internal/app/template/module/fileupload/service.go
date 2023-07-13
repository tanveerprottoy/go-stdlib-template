package fileupload

import (
	"net/http"
	"path/filepath"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/multipart"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/s3pkg"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/uuidpkg"
)

type Service struct {
	s3client *s3pkg.Client
}

func NewService(s3client *s3pkg.Client) *Service {
	s := new(Service)
	s.s3client = s3client
	return s
}

func (s *Service) retrieveSaveFile(key, rootDir, destFileName string, r *http.Request) (string, error) {
	// Retrieve the file from form data
	f, h, err := multipart.GetFormFile(key, r)
	if err != nil {
		return "", err
	}
	return multipart.SaveFile(f, h, rootDir, destFileName+filepath.Ext(h.Filename), r)
}

// handleFilesForKeys saves files on the disk for keys
func (s *Service) handleFilesForKeys(keys []string, rootDir string, destFileNames []string, r *http.Request) ([]string, error) {
	var paths []string
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		f, h, err := multipart.GetFormFile(k, r)
		if err != nil {
			return paths, err
		}
		p, err := multipart.SaveFile(f, h, rootDir, destFileNames[i], r)
		if err != nil {
			return paths, err
		}
		paths = append(paths, p)
	}
	return paths, nil
}

func (s *Service) UploadOne(p []byte, w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	m := map[string]string{"u": ""}
	return m, nil
}

func (s *Service) UploadOneDisk(r *http.Request) (map[string]string, error) {
	m := map[string]string{"path": ""}
	path, err := s.retrieveSaveFile("file", "./uploads", uuidpkg.NewUUIDStr(), r)
	if err != nil {
		return m, err
	}
	m["path"] = path
	return m, nil
}

func (s *Service) UploadMany(r *http.Request) (map[string][]string, error) {
	m := map[string][]string{"u": []string{""}}
	return m, nil
}

func (s *Service) UploadManyDisk(r *http.Request) (map[string][]string, error) {
	m := map[string][]string{"paths": []string{""}}
	return m, nil
}

func (s *Service) UploadManyWithKeysDisk(keys, destFileNames []string, r *http.Request) (map[string][]string, error) {
	m := map[string][]string{"paths": {""}}
	paths, err := s.handleFilesForKeys(keys, "./uploads", destFileNames, r)
	if err != nil {
		return m, err
	}
	m["paths"] = paths
	return m, nil
}
