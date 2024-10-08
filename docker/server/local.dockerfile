FROM golang:alpine


# golang code
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# air for hot reload
RUN go install github.com/air-verse/air@latest

# for http server
EXPOSE 80

CMD ["air", "-c", "server.air.toml"]
