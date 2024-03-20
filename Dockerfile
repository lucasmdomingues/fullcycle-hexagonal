FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install go.uber.org/mock/mockgen@v0.4.0
RUN go install github.com/spf13/cobra-cli@v1.3.0

RUN apt-get update
RUN apt-get install -y sqlite3

CMD ["tail", "-f", "/dev/null"]