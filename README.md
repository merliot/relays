## Relays Device

Relays device controls 1-4 relays using a microcontroller or Rapsberry Pi.

#### Docker

The easiest way to run a relays device is using a pre-built docker image:

```
docker pull ghcr.io/merliot/relays
docker run -p 8000:8000 ghcr.io/merliot/relays
```

Now the device is now running as a web server, listening on port :8000.
