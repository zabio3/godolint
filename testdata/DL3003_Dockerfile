FROM golang:1.12.0-stretch

WORKDIR /go
COPY . /go

RUN cd /usr/src/app && git clone git@github.com:zabio3/godolint.git /usr/src/app

CMD ["go", "run", "main.go"]