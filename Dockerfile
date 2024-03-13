FROM golang:1.21

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

COPY go.mod ./
RUN go mod download

RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install github.com/spf13/cobra-cli@latest

RUN apt-get update
RUN apt-get install -y sqlite3

CMD ["tail", "-f", "/dev/null"]