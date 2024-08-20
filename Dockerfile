FROM golang:1.23-alpine AS builder
WORKDIR /build

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /build/versia-go

FROM alpine:3 AS runner
WORKDIR /app

RUN apk add curl --no-cache

# Copy over some sources to get Sentry's source mapping working in Go
# https://docs.sentry.io/platforms/go/troubleshooting/#missing-stack-trace
COPY ./pkg /app/pkg
COPY ./internal /app/internal
COPY ./ent app/ent

COPY --from=builder /build/versia-go /usr/local/bin/versia-go

ENTRYPOINT [ "/usr/local/bin/versia-go" ]