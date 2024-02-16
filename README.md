## Relays Device

Relays device controls 1-4 relays using a microcontroller or Rapsberry Pi.

#### Docker

The easiest way to run a relays device is using a pre-built docker image.

```
docker pull ghcr.io/merliot/relays/relays
docker run -p 8000:8000 ghcr.io/merliot/relays/relays
```

Now the device is running as a web server, listening on port :8000.  Port :8000 is forwarded to the host so browse to http://\<host\>:8000 (or https://\<host\>:8000) to view and deploy the device.
