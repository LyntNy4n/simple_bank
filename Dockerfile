FROM golang:1.20.5-alpine3.18 as builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY db/migration ./db/migration
COPY start.sh .
COPY wait-for.sh .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
