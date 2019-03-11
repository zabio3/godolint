FROM debian
RUN apt-get update && apt-get clean && rm /var/lib/apt/lists/*

WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]