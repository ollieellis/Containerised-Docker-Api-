FROM golang:1.18-alpine

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./
COPY go.mod /
# RUN go mod download && go mod verify
RUN go mod download

COPY . .
RUN go build -v -o /app

EXPOSE 8080

CMD ["./app"]