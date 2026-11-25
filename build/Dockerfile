FROM golang:1.25.4-trixie AS build
WORKDIR /workdir

COPY go.mod go.sum ./
RUN go mod download

COPY . /workdir
RUN CGO_ENABLED=0 go build -o ./app ./cmd/ticket-service

FROM scratch AS prod
COPY --from=build /workdir/app /app
ENTRYPOINT [ "/app" ]

FROM golang:1.25.4-trixie AS goose
RUN CGO_ENABLED=0 go install github.com/pressly/goose/v3/cmd/goose@v3.27.0

FROM scratch AS migrate
COPY --from=goose /go/bin/goose /goose
COPY migrations /migrations
ENTRYPOINT ["/goose"]



