FROM golang:1.16 AS builder
WORKDIR /usr/share/
COPY ./  ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
WORKDIR /usr/share/
COPY --from=builder /usr/share/app ./
CMD ["./app"] 