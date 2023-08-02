package router

import (
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/router"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/fileupload"

	"github.com/go-chi/chi"
)

func RegisterFileUploadRoutes(router *router.Router, version string, module *fileupload.Module) {
	router.Mux.Route(
		constant.ApiPattern+version+constant.FilesPattern,
		func(r chi.Router) {
			r.Post(constant.RootPattern+"upload-one", module.Handler.UploadOne)
			r.Post(constant.RootPattern+"upload-one-disk", module.Handler.UploadOneDisk)
			r.Post(constant.RootPattern+"upload-many", module.Handler.UploadMany)
			r.Post(constant.RootPattern+"upload-many-disk", module.Handler.UploadManyDisk)
			r.Post(constant.RootPattern+"upload-many-disk-keys", module.Handler.UploadManyWithKeysDisk)
			r.Post(constant.RootPattern+"presigned-one", module.Handler.PutPresignedURLForOne)
		},
	)
}
