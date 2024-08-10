package clients

//
//import (
//	"fmt"
//	"go.opentelemetry.io/otel"
//	"go.opentelemetry.io/otel/attribute"
//	"go.opentelemetry.io/otel/propagation"
//	"go.uber.org/zap"
//	"goOrigin/internal/model/entity"
//	logger2 "goOrigin/pkg/logger"
//	"goOrigin/pkg/utils"
//	"sync"
//	"time"
//)
//
//import (
//	"context"
//	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
//	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
//	"go.opentelemetry.io/otel/log/global"
//	"go.opentelemetry.io/otel/sdk/log"
//	"go.opentelemetry.io/otel/sdk/metric"
//	"go.opentelemetry.io/otel/sdk/resource"
//	"go.opentelemetry.io/otel/sdk/trace"
//	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
//)
//
//func SetupOTelSDK(ctx context.Context, traceInfo map[string]string) (shutdown func(context.Context) error, err error) {
//	var shutdownFuncs []func(context.Context) error
//	shutdown = func(ctx context.Context) error {
//		var err error
//		for _, fn := range shutdownFuncs {
//			fn(ctx)
//		}
//		shutdownFuncs = nil
//		return err
//	}
//
//	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
//	handleErr := func(inErr error) {
//		shutdown(ctx)
//	}
//
//	// Set up propagator.
//	prop := newPropagator()
//	otel.SetTextMapPropagator(prop)
//
//	// Set up trace provider.
//	tracerProvider, err := newTraceProvider(traceInfo)
//	if err != nil {
//		handleErr(err)
//		return
//	}
//	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
//	otel.SetTracerProvider(tracerProvider)
//	// Set up logger provider.
//	loggerProvider, err := newLoggerProvider()
//	if err != nil {
//		handleErr(err)
//		return
//	}
//	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
//	global.SetLoggerProvider(loggerProvider)
//
//	return
//}
//
//type customExporter struct {
//	mu     sync.Mutex
//	traces []entity.TraceEntity
//}
//
//func (e *customExporter) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
//	e.mu.Lock()
//	defer e.mu.Unlock()
//	var (
//		logger, err = logger2.InitZap()
//	)
//	utils.NoError(err)
//	for _, span := range spans {
//		traceID := span.SpanContext().TraceID().String()
//		spanID := span.SpanContext().SpanID().String()
//
//		startTime := span.StartTime().String()
//		duration := span.EndTime().Sub(span.StartTime()).Milliseconds()
//		resources := span.Resource()
//		r, err := span.SpanContext().MarshalJSON()
//		utils.NoError(err)
//		logger.Info(string(r), zap.String("traceID", traceID), zap.String("spanID", spanID), zap.String("startTime", startTime), zap.Int64("duration", duration))
//		traceEntity := entity.TraceEntity{
//			TraceId:   traceID,
//			SpanId:    spanID,
//			SpanKind:  string(span.SpanKind()), // you might need to convert this appropriately
//			Timestamp: startTime,
//			Cost:      fmt.Sprintf("%d", duration),
//			// Add other fields as needed, extracting from attributes or setting default values
//		}
//		for _, attr := range resources.Attributes() {
//			switch attr.Key {
//			case "gid":
//				traceEntity.Gid = attr.Value.AsString()
//			case "cid:":
//				traceEntity.Cid = attr.Value.AsString()
//			case "pid":
//				traceEntity.Pid = attr.Value.AsString()
//			case "az":
//				traceEntity.InstanceZone = attr.Value.AsString()
//			case "app":
//				traceEntity.LocalApp = attr.Value.AsString()
//			case "biz":
//				traceEntity.BusinessId = attr.Value.AsString()
//			case "system":
//				traceEntity.SystemName = attr.Value.AsString()
//			case "sys.baggage":
//				traceEntity.SysBaggage = attr.Value.AsString()
//			case "biz.baggage:":
//				traceEntity.BizBaggage = attr.Value.AsString()
//			}
//		}
//		e.traces = append(e.traces, traceEntity)
//		fmt.Println(utils.ToJson(traceEntity))
//
//	}
//
//	return nil
//}
//
//func (e *customExporter) Shutdown(ctx context.Context) error {
//	return nil
//}
//
//func (e *customExporter) GetTraces() []entity.TraceEntity {
//	e.mu.Lock()
//	defer e.mu.Unlock()
//	return e.traces
//}
//func newPropagator() propagation.TextMapPropagator {
//	return propagation.NewCompositeTextMapPropagator(
//		propagation.TraceContext{},
//		propagation.Baggage{},
//	)
//}
//
//func newTraceProvider(traceInfo map[string]string) (*trace.TracerProvider, error) {
//	var (
//		resourceAttributes []attribute.KeyValue
//	)
//	customExporter := &customExporter{}
//
//	for k, v := range traceInfo {
//		var key, value string
//		key = k
//		value = v
//		resourceAttributes = append(resourceAttributes, attribute.String(key, value))
//	}
//	resourceAttributes = append(resourceAttributes, attribute.String("code", "codd"))
//	traceProvider := trace.NewTracerProvider(
//		//trace.WithBatcher(traceExporter,
//		//	trace.WithBatchTimeout(time.Second)),
//		trace.WithBatcher(customExporter, trace.WithBatchTimeout(time.Second)),
//		trace.WithResource(
//			resource.NewWithAttributes(
//				semconv.SchemaURL,
//				resourceAttributes..., // Add other resource attributes as needed
//			)),
//	)
//	return traceProvider, nil
//}
//
//func newMeterProvider() (*metric.MeterProvider, error) {
//	metricExporter, err := stdoutmetric.New()
//	if err != nil {
//		return nil, err
//	}
//
//	meterProvider := metric.NewMeterProvider(
//		metric.WithReader(metric.NewPeriodicReader(metricExporter,
//			// Default is 1m. Set to 3s for demonstrative purposes.
//			metric.WithInterval(100*time.Second))),
//	)
//	return meterProvider, nil
//}
//
//func newLoggerProvider() (*log.LoggerProvider, error) {
//	logExporter, err := stdoutlog.New()
//	if err != nil {
//		return nil, err
//	}
//
//	loggerProvider := log.NewLoggerProvider(
//		log.WithProcessor(log.NewBatchProcessor(logExporter)),
//	)
//	return loggerProvider, nil
//}
