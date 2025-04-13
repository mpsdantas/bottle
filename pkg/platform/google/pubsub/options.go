package pubsub

import (
	"time"
)

type options struct {
	project    string
	retries    int
	retryDelay time.Duration
}

type OptionFunc = func(option *options)

func WithProject(project string) OptionFunc {
	return func(option *options) {
		option.project = project
	}
}

func WithRetries(retries int) OptionFunc {
	return func(option *options) {
		option.retries = retries
	}
}

func WithRetryDelay(retryDelay time.Duration) OptionFunc {
	return func(option *options) {
		option.retryDelay = retryDelay
	}
}
