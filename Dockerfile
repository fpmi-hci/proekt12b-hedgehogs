FROM golang:alpine

WORKDIR /app

ADD . ./
RUN go mod download
RUN go build -o /app/proekt12b-hedgehogs ./cmd/api


CMD ["./proekt12b-hedgehogs"]
#CMD ["sh", "-c", "tail -f /dev/null"]