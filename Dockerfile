FROM golang:1.22-alpine AS builder

# run the process as an unprivileged user.
RUN mkdir /user

#Install certs
RUN apk add --no-cache ca-certificates

# Working directory outside $GOPATH
WORKDIR /build

# Copy go module files and download dependencies
COPY go.* ./
RUN go mod download

# Copy source files
COPY . .

# Build source
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o go-clinique .

# Minimal image for running the application
FROM scratch as final

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Import the compiled executable from the first stage.
COPY --from=builder ["/build/go-clinique", "/build/.env", "/"]

# Open port
EXPOSE 9000

ENTRYPOINT ["/go-clinique"]
