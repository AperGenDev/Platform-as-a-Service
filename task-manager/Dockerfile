FROM golang:1.26.3 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /docker-my-program .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder ./docker-my-program .       
EXPOSE 8080
CMD [ "./docker-my-program" ]