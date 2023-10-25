// package provider is resposible to provide
// module/component functionalities to other module/component
// used to prevent dependency cycle
package provider

import (
	"context"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/errorext"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user/entity"
)

type User interface {
	ReadOne(id string, ctx context.Context) (entity.User, errorext.HTTPError)
}

type Fileupload interface {
	GetPresignedURLForOne(key string, ctx context.Context) (map[string]string, error)
}
