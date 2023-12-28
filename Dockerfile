# syntax=docker/dockerfile:1

FROM ghcr.io/merliot/device/device-base:latest

WORKDIR /app/relays

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG SCHEME=wss

RUN go work use .
#RUN CGO_ENABLED=0 GOOS=linux go build -tags $SCHEME,prime -o /relays ./cmd/
RUN go build -tags $SCHEME,prime -o /relays ./cmd/
RUN go run ./tools/uf2-builder -target nano-rp2040 -model relays

EXPOSE 8000

ENV PORT_PRIME=8000
CMD ["/relays"]
