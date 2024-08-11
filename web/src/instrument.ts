import * as Sentry from "@sentry/react";
import {ENVIRONMENT, SENTRY_DSN} from "./env.ts";
import {router} from "./routing.ts";

const sampleRate =
    ENVIRONMENT === "development" ? 1 : ENVIRONMENT === "staging" ? 0.5 : 0.2;

Sentry.init({
    dsn: SENTRY_DSN,
    environment: ENVIRONMENT,
    sampleRate,
    tracesSampleRate: 1.0,
    // TODO: Add more targets
    tracePropagationTargets: [/^\//, /^http(s)?:\/\/localhost\//],
    profilesSampleRate: 1.0,
    integrations: [Sentry.tanstackRouterBrowserTracingIntegration(router)],
});

if (SENTRY_DSN) {
    console.log(
        "Sentry initialized with DSN",
        SENTRY_DSN,
        "and environment",
        ENVIRONMENT,
    );
} else {
    console.warn("Sentry not initialized because no DSN is set");
}
