package handlers

import (
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"log"
)

func main() {
	// bootstrap your app...

	// zipkin / opentracing specific stuff
	{
		// set up a span reporter
		reporter := zipkinhttp.NewReporter("http://zipkinhost:9411/api/v2/spans")
		defer reporter.Close()

		// create our local service endpoint
		endpoint, err := zipkin.NewEndpoint("myService", "myservice.mydomain.com:80")
		if err != nil {
			log.Fatalf("unable to create local endpoint: %+v\n", err)
		}

		// initialize our tracer
		nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			log.Fatalf("unable to create tracer: %+v\n", err)
		}

		// use zipkin-go-opentracing to wrap our tracer
		tracer := zipkinot.Wrap(nativeTracer)

		// optionally set as Global OpenTracing tracer instance
		opentracing.SetGlobalTracer(tracer)
	}

	// do other bootstrapping stuff...
}
