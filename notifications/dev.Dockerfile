FROM golang:latest

WORKDIR /app
RUN go get github.com/githubnemo/CompileDaemon
COPY ./go.mod ./
RUN go mod download
COPY ./ ./

ENTRYPOINT CompileDaemon --build="go build -o server src/main.go" --command=./server