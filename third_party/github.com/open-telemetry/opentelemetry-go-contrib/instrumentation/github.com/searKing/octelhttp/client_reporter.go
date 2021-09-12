// Copyright 2021 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package otelgrpc

import (
	"context"
	"net/http"
	"time"

	otelgrpc_ "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
)

type clientReporter struct {
	metrics   *ClientMetrics
	attrs     []attribute.KeyValue
	startTime time.Time
}

func newClientReporter(ctx context.Context, m *ClientMetrics, rpcType grpcType, fullMethod string, target string, source string) *clientReporter {
	r := &clientReporter{
		metrics: m,
	}
	if r.metrics.clientHandledTimeHistogramEnabled {
		r.startTime = time.Now()
	}

	_, attrs := spanInfo(fullMethod, target, source, rpcType)
	r.attrs = attrs
	r.metrics.clientStartedCounter.Add(ctx, 1, r.Attrs()...)
	return r
}

func (r *clientReporter) Attrs(attrs ...attribute.KeyValue) []attribute.KeyValue {
	attrs = append(r.attrs, attrs...)
	filter := AttrsFilter
	if filter != nil {
		return filter(attrs...)
	}
	return attrs
}

func (r *clientReporter) ReceivedResponse(ctx context.Context, resp *http.Response) {
	attrs := r.Attrs(otelgrpc_.RPCMessageTypeReceived)
	r.metrics.clientStreamRequestReceived.Add(ctx, 1, attrs...)
	if r.metrics.clientStreamReceiveSizeHistogramEnabled {
		if resp != nil {
			r.metrics.clientStreamReceiveSizeHistogram.Record(ctx, resp.ContentLength, attrs...)
		} else {
			r.metrics.clientStreamReceiveSizeHistogram.Record(ctx, -1, attrs...)
		}
	}
}

func (r *clientReporter) SentRequest(ctx context.Context, req *http.Request) {
	attrs := r.Attrs(otelgrpc_.RPCMessageTypeSent)
	r.metrics.clientStreamRequestSent.Add(ctx, 1, attrs...)
	if r.metrics.clientStreamSendSizeHistogramEnabled {
		if req != nil {
			r.metrics.clientStreamSendSizeHistogram.Record(ctx, req.ContentLength, attrs...)
		} else {
			r.metrics.clientStreamSendSizeHistogram.Record(ctx, -1, attrs...)
		}
	}
}

func (r *clientReporter) Handled(ctx context.Context, code codes.Code) {
	attrs := r.Attrs(statusCodeAttr(code))
	r.metrics.clientHandledCounter.Add(ctx, 1, attrs...)

	if r.metrics.clientHandledTimeHistogramEnabled {
		r.metrics.clientHandledTimeHistogram.Record(ctx, time.Since(r.startTime).Seconds(), attrs...)
	}
}
