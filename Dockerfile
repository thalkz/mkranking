FROM golang:1.18-alpine as builder
WORKDIR app/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /server

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY ./templates /templates
COPY ./static /static
COPY --from=builder /server /server
ENTRYPOINT ["/server"]