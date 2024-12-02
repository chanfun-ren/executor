// pkg/interceptor/interceptors.go
package interceptor

import (
	"context"
	"time"

	"github.com/chanfun-ren/executor/pkg/logging"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

var (
	// 定义一个全局的 HistogramVec 指标
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_seconds",
			Help:    "Duration of gRPC requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func init() {
	prometheus.MustRegister(requestDuration)
}

func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log := logging.FromContext(ctx)
	log.Infow("Received request", "method", info.FullMethod, "request", req)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Errorw("Failed to handle request", "method", info.FullMethod, "error", err)
	}
	log.Infow("Handled request", "method", info.FullMethod, "response", resp)
	return resp, err
}

func MetricsInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(start).Seconds()

	requestDuration.WithLabelValues(info.FullMethod).Observe(duration)

	return resp, err
}