FROM golang:1.20

COPY . ./src

WORKDIR ./src

RUN go get

RUN go build main.go

EXPOSE 8080

CMD ["./main"]