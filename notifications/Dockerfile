FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build ./src

FROM scratch
WORKDIR /app
COPY --from=build /app/build ./
CMD ["./build"]