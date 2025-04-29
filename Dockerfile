FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/mosmetro ./cmd/mosmetro

FROM alpine 

WORKDIR /app

COPY --from=builder /app/mosmetro /app/mosmetro

CMD [ "./mosmetro" ]