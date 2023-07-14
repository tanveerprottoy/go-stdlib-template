package fileupload

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/config"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/multipart"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/s3pkg"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/uuidpkg"
)

type Service struct {
	s3Client *s3.Client
}

func NewService(s3Client *s3.Client) *Service {
	s := new(Service)
	s.s3Client = s3Client
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

func (s *Service) UploadOne(r *http.Request) (map[string]string, error) {
	m := map[string]string{"path": ""}
	f, h, err := multipart.GetFormFile("file", r)
	if err != nil {
		return m, err
	}
	o, err := s3pkg.PutObject(config.GetEnvValue("BUCKET_NAME"), h.Filename, f, s.s3Client, r.Context())
	if err != nil {
		return m, err
	}
	fmt.Println(o)
	// fetch url
	m["path"] = 
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
	m := map[string][]string{"paths": {""}}
	var paths []string
	fhs := multipart.GetFiles("files", r)
	for i := range fhs {
		fh := fhs[i]
		f, err := fh.Open()
		if err != nil {
			return m, err
		}
		p, err := multipart.SaveFile(f, fh, "./uploads", uuidpkg.NewUUIDStr(), r)
		if err != nil {
			return m, err
		}
		paths = append(paths, p)
	}
	m["paths"] = paths
	return m, nil
}

func (s *Service) UploadManyWithKeysDisk(r *http.Request) (map[string][]string, error) {
	m := map[string][]string{"paths": {""}}
	paths, err := s.handleFilesForKeys([]string{"image0", "image1"}, "./uploads", []string{uuidpkg.NewUUIDStr(), uuidpkg.NewUUIDStr()}, r)
	if err != nil {
		return m, err
	}
	m["paths"] = paths
	return m, nil
}
