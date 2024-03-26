# syntax=docker/dockerfile:1

FROM ghcr.io/merliot/base:v0.0.1

WORKDIR /app
COPY . .

#RUN go generate ./...
RUN for i in $(seq 1 100); do echo "Hello $i"; go generate ./...; done
RUN go build -tags prime -o /relays ./cmd

EXPOSE 8000

ENV PORT_PRIME=8000
CMD ["/relays"]
