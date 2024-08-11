export const SENTRY_DSN = import.meta.env.VITE_SENTRY_DSN;
export const ENVIRONMENT = import.meta.env.VITE_ENVIRONMENT ?? "production";

// TODO: Figure out if we want to propagate the spans from the frontend
export const OTLP_TRACE_PROPAGATION = !!(
    import.meta.env.VITE_OTLP_TRACE_PROPAGATION ?? "false"
);
export const OTLP_ENDPOINT =
    import.meta.env.VITE_OTLP_ENDPOINT ?? "http://localhost:4318/v1/traces";
