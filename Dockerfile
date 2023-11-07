FROM golang:1.21
ENV ROOT=/go/src/app
RUN apt-get upgrade && apt-get install -y make
WORKDIR ${ROOT}

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]
