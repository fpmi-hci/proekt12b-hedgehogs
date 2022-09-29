FROM golang:alpine

WORKDIR /app

ADD . ./
RUN go mod download
RUN go build -o /app/project ./cmd/api


CMD ["./project"]
#CMD ["sh", "-c", "tail -f /dev/null"]