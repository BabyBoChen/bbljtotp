FROM golang:1.22.3

COPY . ./src

WORKDIR ./src

RUN go get

RUN go build main.go

EXPOSE 80

CMD ["./main"]