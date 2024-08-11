import {layer} from "@effect/opentelemetry/WebSdk";
import {OTLPTraceExporter} from "@opentelemetry/exporter-trace-otlp-http";
import {BatchSpanProcessor} from "@opentelemetry/sdk-trace-base";
import {Layer, ManagedRuntime} from "effect";
import {OTLP_ENDPOINT, OTLP_TRACE_PROPAGATION} from "../env.ts";

export const WebSdkLive = OTLP_TRACE_PROPAGATION
    ? layer(() => ({
        resource: {serviceName: "Web UI"},
        spanProcessor: new BatchSpanProcessor(
            new OTLPTraceExporter({url: OTLP_ENDPOINT}),
        ),
    }))
    : Layer.empty;

if (OTLP_ENDPOINT) {
    console.log("OpenTelemetry initialized with OTLP endpoint", OTLP_ENDPOINT);

    if (OTLP_TRACE_PROPAGATION) {
        console.log("OpenTelemetry initialized with trace propagation");
    } else {
        console.warn("OpenTelemetry initialized without trace propagation");
    }
} else {
    console.warn("OpenTelemetry not initialized because no OTLP endpoint is set");
}

export const runtime = ManagedRuntime.make(WebSdkLive);
