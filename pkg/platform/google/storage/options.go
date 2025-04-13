package storage

import (
	"cloud.google.com/go/storage"
)

type WriterOption func(*storage.Writer)

func WithCacheControl(v string) WriterOption {
	return func(w *storage.Writer) {
		w.CacheControl = v
	}
}

func WithContentType(v string) WriterOption {
	return func(w *storage.Writer) {
		w.ContentType = v
	}
}

func WithContentEncoding(v string) WriterOption {
	return func(w *storage.Writer) {
		w.ContentEncoding = v
	}
}

func WithContentDisposition(v string) WriterOption {
	return func(w *storage.Writer) {
		w.ContentDisposition = v
	}
}

func WithMetadata(m map[string]string) WriterOption {
	return func(w *storage.Writer) {
		w.Metadata = m
	}
}
