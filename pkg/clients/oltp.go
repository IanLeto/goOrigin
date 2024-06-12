package clients

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"goOrigin/internal/model/entity"
	"goOrigin/pkg/utils"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

type customExporter struct {
	mu     sync.Mutex
	traces []entity.TraceEntity
}

func (e *customExporter) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	for _, span := range spans {
		traceID := span.SpanContext().TraceID().String()
		spanID := span.SpanContext().SpanID().String()
		//operation := span.Name()
		startTime := span.StartTime().String()
		//endTime := span.EndTime().String()
		duration := span.EndTime().Sub(span.StartTime()).Milliseconds()
		attributes := span.Attributes()
		r, err := span.SpanContext().MarshalJSON()
		utils.NoError(err)
		fmt.Println(string(r))
		traceEntity := entity.TraceEntity{
			TraceId:   traceID,
			SpanId:    spanID,
			SpanKind:  string(span.SpanKind()), // you might need to convert this appropriately
			Timestamp: startTime,
			Cost:      fmt.Sprintf("%d", duration),
			// Add other fields as needed, extracting from attributes or setting default values
		}
		fmt.Println(utils.ToJson(attributes))
		for _, attr := range attributes {
			switch attr.Key {
			case "db.system":
				traceEntity.SystemName = attr.Value.AsString()
			case "http.method":
				traceEntity.ReqMethod = attr.Value.AsString()
			case "http.url":
				traceEntity.ReqUrl = attr.Value.AsString()
			case "net.peer.name":
				traceEntity.RemoteHost = attr.Value.AsString()
			case "net.peer.port":
				traceEntity.RemotePort = attr.Value.AsString()
			case "http.status_code":
				traceEntity.ResultCode = attr.Value.AsString()
				// Add other cases to map attributes to TraceEntity fields
			}
			e.traces = append(e.traces, traceEntity)
		}
	}

	return nil
}

func (e *customExporter) Shutdown(ctx context.Context) error {
	return nil
}

func (e *customExporter) GetTraces() []entity.TraceEntity {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.traces
}

func SetupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			fn(ctx)
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		shutdown(ctx)
	}

	// Set up propagator.
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// Set up trace provider.
	tracerProvider, err := newTraceProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	// Set up meter provider.
	meterProvider, err := newMeterProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	// Set up logger provider.
	loggerProvider, err := newLoggerProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)

	return
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider() (*trace.TracerProvider, error) {
	//traceExporter, err := stdouttrace.New(
	//	stdouttrace.WithPrettyPrint())
	//if err != nil {
	//	return nil, err
	//}
	customExporter := &customExporter{}

	traceProvider := trace.NewTracerProvider(
		//trace.WithBatcher(traceExporter,
		//	trace.WithBatchTimeout(time.Second)),
		trace.WithBatcher(customExporter,
			trace.WithBatchTimeout(time.Second)),

		trace.WithResource(resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceNameKey.String("ian"),
			attribute.String("test", "xxx"),
		)),
	)
	return traceProvider, nil
}

func newMeterProvider() (*metric.MeterProvider, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			metric.WithInterval(100*time.Second))),
	)
	return meterProvider, nil
}

func newLoggerProvider() (*log.LoggerProvider, error) {
	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)
	return loggerProvider, nil
}
