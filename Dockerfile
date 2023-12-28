# syntax=docker/dockerfile:1

FROM ghcr.io/merliot/device/device-base:latest

WORKDIR /app/relays

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG SCHEME=wss

RUN go work use .
RUN CGO_ENABLED=0 GOOS=linux go build -tags $SCHEME,prime -o /relays ./cmd/
RUN tinygo build -target nano-rp2040 -o nano-rp2040.uf2 -size short -stack-size 8kb ./run/

EXPOSE 8000

ENV PORT_PRIME=8000
CMD ["/relays"]
