package s3pkg

import (
	"mime/multipart"
)

// S3Object type creates a union of expected types
type S3Object interface {
	multipart.File
}