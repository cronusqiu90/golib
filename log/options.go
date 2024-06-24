package log

import "io"

type Option interface {
	apply(*Logger)
}

type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
	f(log)
}

func WithOutput(w io.Writer) Option {
	return optionFunc(func(log *Logger) {
		log.out = w
	})
}

func WithLevel(l Level) Option {
	return optionFunc(func(log *Logger) {
		log.level = l
	})
}

func WithEnableCaller() Option {
	return optionFunc(func(log *Logger) {
		log.callerEnabled = true
	})
}

func AddCallerSkip(skip int) Option {
	return optionFunc(func(log *Logger) {
		log.callerSkip = skip
	})
}
