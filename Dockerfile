FROM golang:latest

RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    make \
 && rm -rf /var/lib/apt/lists/*

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest

WORKDIR /app

EXPOSE 40000
EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]
