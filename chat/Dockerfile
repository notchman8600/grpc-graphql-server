FROM golang:latest as builder
WORKDIR /workdir
COPY . .

# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go generate .
RUN go get
RUN go build -o chat

# Runtime Container
FROM gcr.io/distroless/base
COPY --from=builder /workdir/chat /chat
ENTRYPOINT ["/chat"]
