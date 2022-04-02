FROM golang:latest as builder
WORKDIR /workdir
ENV GOPATH=/workdir
ENV GOBIN=/workdir
RUN go install github.com/movio/bramble/cmd/bramble@latest

FROM gcr.io/distroless/base
COPY --from=builder /workdir/bramble /bramble
COPY config.json /config.json
ENTRYPOINT ["/bramble","-conf","/config.json"]
