FROM golang:1.17-alpine

WORKDIR /go/src

COPY . .

RUN go build -o /bin/metrics main.go

RUN chmod +x /bin/metrics

EXPOSE 8181

ENTRYPOINT [ "metrics" ]