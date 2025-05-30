FROM golang:1.24-alpine3.20 as builder
COPY . /src
WORKDIR /src/cmd
RUN apk add --no-cache git ca-certificates && \
  go build -v .

FROM alpine:3.22
COPY --from=builder /src/cmd /bin/godolint
ENTRYPOINT ["/bin/godolint"]
