FROM golang:1.23-alpine3.20 as builder
COPY . /src
WORKDIR /src/cmd
RUN apk add --no-cache git ca-certificates && \
  go build -v .

FROM alpine:3.20
COPY --from=builder /src/cmd /bin/godolint
ENTRYPOINT ["/bin/godolint"]
