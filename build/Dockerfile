FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .
#################
FROM scratch

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]