package codestyle

import (
	"fmt"
	"go.uber.org/zap"
)

type IService interface {
	send(msg string)
	recv() string
}

type Service struct {
}

func (s *Service) send(msg string) {
	fmt.Println("send:", msg)
}

func (s *Service) recv() string {
	return "recv"
}

type ServiceMiddleware func(IService) IService

type loggingMiddleware struct {
	IService
	logger *zap.Logger
}

func LoggingMiddleware(logger *zap.Logger) ServiceMiddleware {
	return func(next IService) IService {
		return &loggingMiddleware{
			IService: next,
			logger:   logger,
		}
	}
}

func (s *loggingMiddleware) send(msg string) {
	s.logger.Info("log send")
	defer s.logger.Info("log send ok")
	s.IService.send(msg)
}

func (s *loggingMiddleware) recv() string {
	s.logger.Info("log recv")
	defer s.logger.Info("log recv ok")
	return s.IService.recv()
}

type Metrics int

type metricsMiddleware struct {
	IService
	metricCount *Metrics
}

func MetricsMiddleware(metricCount *Metrics) ServiceMiddleware {
	return func(next IService) IService {
		return &metricsMiddleware{
			IService:    next,
			metricCount: metricCount,
		}
	}
}

func (m *metricsMiddleware) send(msg string) {
	*m.metricCount++
	m.IService.send(msg)
}

func (m *metricsMiddleware) recv() string {
	fmt.Println("metricCount:", *m.metricCount)
	return m.IService.recv()
}

func MiddlewareUsage() {
	var s IService
	s = &Service{}
	logger := zap.NewExample()
	s = LoggingMiddleware(logger)(s)
	metrics := Metrics(0)
	s = MetricsMiddleware(&metrics)(s)
	s.send("hello")
	s.recv()
}
