package codestyle

import (
	"fmt"
	"go.uber.org/zap"
)

type IService interface {
	send(msg string)
	receive() string
}

type service struct {
}

func (s service) send(msg string) {
	fmt.Println("send:", msg)
}

func (s service) receive() string {
	return "received"
}

type ServiceMiddleware func(IService) IService

type loggingMiddleware struct {
	IService
	logger *zap.Logger
}

func LoggingMiddleware(logger *zap.Logger) ServiceMiddleware {
	return func(next IService) IService {
		return loggingMiddleware{
			IService: next,
			logger:   logger,
		}
	}
}

func (s loggingMiddleware) send(msg string) {
	s.logger.Info("log send")
	defer s.logger.Info("log send ok")
	s.IService.send(msg)
}

func (s loggingMiddleware) receive() string {
	s.logger.Info("log receive")
	defer s.logger.Info("log receive ok")
	return s.IService.receive()
}

type Metrics int

type metricsMiddleware struct {
	IService
	metricCount *Metrics
}

func MetricsMiddleware(metricCount *Metrics) ServiceMiddleware {
	return func(next IService) IService {
		return metricsMiddleware{
			IService:    next,
			metricCount: metricCount,
		}
	}
}

func (m metricsMiddleware) send(msg string) {
	*m.metricCount++
	m.IService.send(msg)
}

func (m metricsMiddleware) receive() string {
	fmt.Println("metricCount:", *m.metricCount)
	return m.IService.receive()
}

func MiddlewareUsage() {
	var s IService
	s = service{}
	logger := zap.NewExample()
	s = LoggingMiddleware(logger)(s)
	metrics := Metrics(0)
	s = MetricsMiddleware(&metrics)(s)
	s.send("hello")
	s.receive()
}
