FROM golang:1.17-alpine as builder
WORKDIR app/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /server

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /server /server
EXPOSE 8080 8080
ENTRYPOINT ["/server"]