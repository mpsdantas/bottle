package bottle

type options struct {
	uploadLimit int64
}

type Option func(*options)

func WithUploadLimit(limit int64) Option {
	return func(o *options) {
		o.uploadLimit = limit
	}
}
