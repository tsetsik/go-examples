FROM golang:1.24-alpine3.22

# Go build env settings
ENV GOOS="linux"
ENV GOARCH="amd64"
ENV CGO_ENABLED=0

WORKDIR /app

# Cache go modules
COPY go.mod ./
RUN go mod download

# Copy project files and build 
COPY . .
RUN go build -o ./bin/job-queue ./cmd/service/service.go

EXPOSE 9393

ENTRYPOINT [ "/app/bin/job-queue" ]