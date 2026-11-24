FROM golang:1.25.4-trixie AS build
COPY . /workdir
RUN go build -o /workdir/cmd/ticket-service/main.go /workdir/app

FROM scratch AS prod
COPY --from=build /workdir/app app
ENTRYPOINT [ "app" ]

FROM golang:1.25.4-trixie AS goose
RUN CGO_ENABLED=0 go install github.com/pressly/goose/v3/cmd/goose@v3.27.0

FROM scratch AS migrate
COPY --from=goose /go/bin/goose /goose
COPY migrations /migrations
ENTRYPOINT ["/goose"]



