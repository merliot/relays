# syntax=docker/dockerfile:1

#FROM tinygo/tinygo-dev:latest
FROM ghcr.io/merliot/hub/hub-base:latest

WORKDIR /app/relays

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG SCHEME=https

RUN go work init
RUN go work use .
RUN CGO_ENABLED=0 GOOS=linux go build -tags $SCHEME,prime -o /relays ./cmd/

EXPOSE 8000

ENV PORT_PRIME=8000
CMD ["/relays"]
