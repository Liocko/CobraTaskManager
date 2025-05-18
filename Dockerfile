FROM golang:1.23
LABEL authors="liocko"
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/cobraTaskManager main.go
CMD ["/app/cobraTaskManager"]