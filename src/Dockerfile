#Stage 1: Build
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN CGO_ENABLED=0 go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
ENV PATH=$PATH:/go/bin

COPY . .
RUN CGO_ENABLED=0 GOOD=linux go build -o server ./main.go

#Stage 2: Run
FROM alpine:3.18

WORKDIR /root/

RUN apk add --no-cache postgresql-client
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/database/migrations ./database/migrations
COPY --from=builder /app/server .
COPY --from=builder /app/scripts ./scripts
RUN chmod +x ./scripts/migrate.sh

EXPOSE 5000

CMD [ "./server" ]

