FROM golang:1.17-alpine as build

WORKDIR app/

COPY go.mod ./
RUN go mod download

COPY src/ ./
RUN GOOS=linux GOARCH=amd64 go build -o /server

FROM scratch
WORKDIR /
COPY --from=build /server /server
EXPOSE 8080
CMD ["/server"]