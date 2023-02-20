FROM golang:1.20-alpine AS builder
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix 'static' -o app .
RUN chmod +x app

FROM golang:1.20-alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /src/app .
COPY --from=builder /src/.env .
EXPOSE 8000
CMD ["./app"]
