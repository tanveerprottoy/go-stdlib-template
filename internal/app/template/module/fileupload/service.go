package fileupload

import (
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/multipart"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/s3pkg"
)

type Service struct {
	s3client *s3pkg.Client
}

func NewService(s3client *s3pkg.Client) *Service {
	s := new(Service)
	s.s3client = s3client
	return s
}

func (s *Service) UploadOne(p []byte, w http.ResponseWriter, r *http.Request) (map[string]string, error) {

	/* 	defer f.Close()
	   	p, err = file.SaveFile(f, rootDir, header.Filename)
	   	if err != nil {
	   		return paths, err
	   	}
	   	response.Respond(http.StatusOK, map[string]string{"url": urls[0]}, w) */
	m := map[string]string{"u": ""}
	return m, nil
}

func (s *Service) UploadOneDisk(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	m := map[string]string{"url": ""}
	urls, err := multipart.HandleFiles(r, []string{"file"}, "./uploads")
	if err != nil {
		return m, err
	}
	m["url"] = urls[0]
	return m, nil
}

func (s *Service) UploadMany(w http.ResponseWriter, r *http.Request) {

}

func (s *Service) UploadManyDisk(w http.ResponseWriter, r *http.Request) {

}
