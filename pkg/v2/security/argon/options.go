package argon

type options struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type Option func(*options)

func WithMemory(memory uint32) Option {
	return func(o *options) {
		o.memory = memory
	}
}

func WithIterations(iterations uint32) Option {
	return func(o *options) {
		o.iterations = iterations
	}
}

func WithParallelism(parallelism uint8) Option {
	return func(o *options) {
		o.parallelism = parallelism
	}
}

func WithSaltLength(saltLength uint32) Option {
	return func(o *options) {
		o.saltLength = saltLength
	}
}

func WithKeyLength(keyLength uint32) Option {
	return func(o *options) {
		o.keyLength = keyLength
	}
}
