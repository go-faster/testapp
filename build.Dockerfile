FROM golang:1.23


WORKDIR /tmp/cache
COPY go.mod go.sum /tmp/cache/
RUN go mod download

COPY . /testapp
WORKDIR /testapp
RUN go build -o testapp ./cmd/testapp
RUN mv testapp /usr/bin/testapp

ENTRYPOINT ["/usr/bin/testapp"]
