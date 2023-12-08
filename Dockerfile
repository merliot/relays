# syntax=docker/dockerfile:1

FROM tinygo/tinygo-dev:latest

WORKDIR /app/relays

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG SCHEME=http

RUN go work init
RUN go work use .
RUN CGO_ENABLED=0 GOOS=linux go build -tags $SCHEME,prime -o /relays ./cmd/

EXPOSE 8000

ENV PORT_PRIME=8000
CMD ["/relays"]
