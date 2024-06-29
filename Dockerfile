FROM golang:1.18-alpine as builder
COPY . /src
WORKDIR /src/cmd
RUN apk add --no-cache git ca-certificates && \
  go build -v .

FROM alpine:3.20.1
COPY --from=builder /src/cmd /bin/godolint
ENTRYPOINT ["/bin/godolint"]