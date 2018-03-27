FROM golang:1.8

RUN mkdir -p /go/src/github.com/micahlee/big-earl-cli/output
WORKDIR /go/src/github.com/micahlee/big-earl-cli

COPY . .

RUN go get

ENV GOOS=linux
ENV GOARCH=amd64

ENTRYPOINT ["/usr/local/go/bin/go"]
CMD ["build", "-v"]
