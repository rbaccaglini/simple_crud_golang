FROM amd64/golang:1.23 AS BUILDER

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/simple_crud_golang

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/meuprimeirocrudgo .

FROM golang:1.23-alpine AS RUNNER

WORKDIR /root/

COPY --from=BUILDER /app/meuprimeirocrudgo .

EXPOSE 8080

CMD ["./meuprimeirocrudgo"]
