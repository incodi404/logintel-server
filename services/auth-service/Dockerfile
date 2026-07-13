FROM golang:1.25-alpine as base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM base as dev 
CMD [ "go", "run", "main.go" ]