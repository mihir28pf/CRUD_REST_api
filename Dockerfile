FROM golang:alpine as builder01
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main1 .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder01 /app/main .
COPY --from=builder01 /app/config.json .
EXPOSE 8080
CMD ["./main1"]
