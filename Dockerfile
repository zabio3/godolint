FROM golang:1.16.5-alpine3.12 as builder
COPY . /src
WORKDIR /src/cmd
RUN apk add --no-cache git ca-certificates && \
  go build -v .

FROM alpine:3.12.0
COPY --from=builder /src/cmd /bin/godolint
ENTRYPOINT ["/bin/godolint"]