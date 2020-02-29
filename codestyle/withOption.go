package codestyle

import (
	"fmt"
	"go.uber.org/zap"
)

type option interface {
	apply(*options)
}

type options struct {
	tls    bool
	logger *zap.Logger
}

type tlsOption bool

func (t *tlsOption) apply(opt *options) {
	opt.tls = bool(*t)
}

func WithTls(t bool) option {
	tt := tlsOption(t)
	return &tt
}

type loggerOption struct {
	Log *zap.Logger
}

func (l *loggerOption) apply(opt *options) {
	opt.logger = l.Log
}

func WithLogger(log *zap.Logger) option {
	return &loggerOption{Log: log}
}

func Send(endpoint string, opt ...option) {
	options := options{
		tls:    false,
		logger: nil,
	}
	for _, o := range opt {
		o.apply(&options)
	}
	if options.tls == true {
		fmt.Println("tls opened")
	}
	if options.logger != nil {
		fmt.Println("log opened")
	}
	fmt.Println("send to:", endpoint)
}

func WithOptionUsage() {
	Send("192.168.11.1:8080")
	Send("192.168.11.1:8080", WithTls(true))
	Send("192.168.11.1:8080", WithLogger(zap.NewExample()))
}
