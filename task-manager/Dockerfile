FROM golang:1.26.3

WORKDIR /app

COPY task-manager/go.mod task-manager/go.sum ./
RUN go mod download

COPY task-manager/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-my-program

EXPOSE 8080

CMD [ "/docker-my-program" ]