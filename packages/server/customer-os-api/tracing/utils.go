package tracing

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

type spanCtxKey struct{}

var activeSpanCtxKey = spanCtxKey{}

const (
	SpanTagTenant    = "tenant"
	SpanTagUserId    = "user-id"
	SpanTagComponent = "component"
)

func ExtractSpanCtx(ctx context.Context) opentracing.SpanContext {
	if ctx.Value(activeSpanCtxKey) != nil {
		return ctx.Value(activeSpanCtxKey).(opentracing.SpanContext)
	}
	return nil
}

func EnrichCtxWithSpanCtxForGraphQL(ctx context.Context, operationContext *graphql.OperationContext) context.Context {
	spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(operationContext.Headers))
	if err != nil {
		return ctx
	}
	if ExtractSpanCtx(ctx) != nil {
		return ctx
	}
	return context.WithValue(ctx, activeSpanCtxKey, spanCtx)
}

func StartGraphQLTracerSpan(ctx context.Context, operationName string, operationContext *graphql.OperationContext) (context.Context, opentracing.Span) {
	spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(operationContext.Headers))

	if err != nil {
		rootSpan := opentracing.GlobalTracer().StartSpan(operationName)
		opentracing.GlobalTracer().Inject(rootSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(operationContext.Headers))
		return opentracing.ContextWithSpan(ctx, rootSpan), rootSpan
	}

	serverSpan := opentracing.GlobalTracer().StartSpan(operationName, ext.RPCServerOption(spanCtx))
	return opentracing.ContextWithSpan(ctx, serverSpan), serverSpan
}

func StartHttpServerTracerSpanWithHeader(ctx context.Context, operationName string, headers http.Header) (context.Context, opentracing.Span) {
	spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(headers))

	if err != nil {
		serverSpan := opentracing.GlobalTracer().StartSpan(operationName)
		opentracing.GlobalTracer().Inject(serverSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(headers))
		return opentracing.ContextWithSpan(ctx, serverSpan), serverSpan
	}

	serverSpan := opentracing.GlobalTracer().StartSpan(operationName, ext.RPCServerOption(spanCtx))
	return opentracing.ContextWithSpan(ctx, serverSpan), serverSpan
}

func TraceErr(span opentracing.Span, err error) {
	tracing.TraceErr(span, err)
}

func setDefaultSpanTags(ctx context.Context, span opentracing.Span) {
	if common.GetTenantFromContext(ctx) != "" {
		span.SetTag(SpanTagTenant, common.GetTenantFromContext(ctx))
	}
	if common.GetUserIdFromContext(ctx) != "" {
		span.SetTag(SpanTagUserId, common.GetUserIdFromContext(ctx))
	}
}

func SetDefaultResolverSpanTags(ctx context.Context, span opentracing.Span) {
	setDefaultSpanTags(ctx, span)
	span.SetTag(SpanTagComponent, constants.ComponentResolver)
}

func SetDefaultServiceSpanTags(ctx context.Context, span opentracing.Span) {
	setDefaultSpanTags(ctx, span)
	span.SetTag(SpanTagComponent, constants.ComponentService)
}

func SetDefaultNeo4jRepositorySpanTags(ctx context.Context, span opentracing.Span) {
	setDefaultSpanTags(ctx, span)
	span.SetTag(SpanTagComponent, constants.ComponentNeo4jRepository)
}
