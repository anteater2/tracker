FROM golang:1.9.1

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/anteater2/api

COPY . .

RUN dep ensure -vendor-only
RUN go-wrapper install

EXPOSE 5000

CMD ["go-wrapper", "run"]