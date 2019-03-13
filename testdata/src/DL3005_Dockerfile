FROM debian:1.12.0-stretch
RUN apt-get update && apt-get upgrade && apt-get clean && rm /var/lib/apt/lists/*

WORKDIR /go
COPY . /go

CMD ["go", "run", "main.go"]