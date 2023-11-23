FROM golang:1.21-alpine as builder
COPY . /src
WORKDIR /src/cmd
RUN apk add --no-cache git ca-certificates && \
  go build -v .

FROM alpine:3.18.4
COPY --from=builder /src/cmd /bin/godolint
ENTRYPOINT ["/bin/godolint"]