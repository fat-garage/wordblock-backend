FROM golang:1.18 as golang-builder
WORKDIR /go/src/github.com/fat-garage/wordblock-backend
COPY . .
COPY main.go ./main.go
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=golang-builder /go/src/github.com/fat-garage/wordblock-backend/app .
COPY --from=golang-builder /go/src/github.com/fat-garage/wordblock-backend/docs ./docs
EXPOSE 8080
ENTRYPOINT ["./app"]