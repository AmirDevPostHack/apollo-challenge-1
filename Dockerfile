FROM golang:1.21 as build

LABEL maintainer="amir.academicinquiry@gmail.com"
LABEL description="proxy server for Github Gists REST API"

WORKDIR /go/src/app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/server ./server.go

FROM gcr.io/distroless/static-debian12

COPY --from=build /go/bin/server /

EXPOSE 8080
CMD ["/server"]
