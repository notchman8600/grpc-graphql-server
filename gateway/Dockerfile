FROM golang:latest as builder
WORKDIR /workdir
COPY . .

# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go build -o gateway

# Runtime Container
FROM gcr.io/distroless/base
COPY --from=builder /workdir/gateway /gateway
ENTRYPOINT ["/gateway"]
