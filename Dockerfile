FROM golang:1.9.1

WORKDIR /go/src/app

COPY . .

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 5000

CMD ["go-wrapper", "run"]