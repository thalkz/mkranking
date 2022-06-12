FROM golang:1.18-alpine as builder
WORKDIR app/
COPY ./server/go.mod ./
COPY ./server/go.sum ./
RUN go mod download
COPY ./server ./
RUN go build -o /main

FROM alpine as runner
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY ./templates /templates
COPY ./static /static
COPY --from=builder /main /server/main
ENTRYPOINT ["/server/main"]