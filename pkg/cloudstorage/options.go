package cloudstorage

import (
	"io"
)

type UploadOptions struct {
	Filename     string
	Data         io.Reader
	CacheControl string
}
