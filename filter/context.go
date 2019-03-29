package filter

import (
	"context"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/metadata"
	"github.com/sirupsen/logrus"
	"gopkg.in/fatih/set.v0"
)

const (
	prefix           = "x-weather-"
	TraceIDKey       = "trace_id"
	traceIDHeaderKey = "x-weather-trace-id"
)

// ContextBuilder creates context objects from http request headers
type ContextBuilder struct {
	logger   *logrus.Entry
	required set.Interface
}

// NewContextBuilder creates a new context builder
func NewContextBuilder(logger *logrus.Entry) (*ContextBuilder, error) {
	builder := &ContextBuilder{
		logger: logger,
	}

	return builder, nil
}

// CreateContext appends context objects to request object using headers
func (b *ContextBuilder) CreateContext(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	ctx := context.Background()

	// add trace id to context metadata for downstream logging methods
	traceID := req.Request.Header.Get(traceIDHeaderKey)
	if traceID != "" {
		//Preserve metadata
		md, ok := metadata.FromContext(ctx)
		if !ok {
			md = metadata.Metadata{
				TraceIDKey: traceID,
			}
		} else {
			md[TraceIDKey] = traceID
		}
		// Set latest metadata to context
		ctx = metadata.NewContext(ctx, md)
	}

	req.SetAttribute("ctx", ctx)

	chain.ProcessFilter(req, resp)
}
