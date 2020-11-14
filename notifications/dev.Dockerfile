FROM golang:alpine

WORKDIR /app
COPY ./go.mod ./
RUN go mod download
COPY ./ ./

ENTRYPOINT ["go", "run", "./src/main.go"]