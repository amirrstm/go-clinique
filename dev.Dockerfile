# DO NOT use in production!
FROM golang:1.22-alpine

# change working dir
WORKDIR /app

# Copy go module files and download dependencies
COPY go.* ./
RUN go mod download

# install file watcher
RUN apk add make && go install github.com/githubnemo/CompileDaemon && go install github.com/swaggo/swag/cmd/swag

ENTRYPOINT ["CompileDaemon", "-exclude-dir=.git", "-exclude-dir=docs", "-build=make build", "-command=./build/go-clinique", "-graceful-kill=true"]
