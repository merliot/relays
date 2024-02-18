## Relays Device

Relays device controls 1-4 relays using a microcontroller or Rapsberry Pi.

#### Docker

The easiest way to run a relays device is using a pre-built docker image:

```
docker pull ghcr.io/merliot/relays
docker run -p 8000:8000 ghcr.io/merliot/relays
```

Now the device is now running in a docker container as a web server, listening on port :8000.  

Browse to http://\<host\>:8000 to view and deploy the device, where \<host\> is the name or IP address of host running the container.

If the docker host is using https://, pass in the environment variable WS_SCHEME=wss://.

```
docker run -e "WS_SCHEME=wss://" -p 8000:8000 ghcr.io/merliot/relays
```

#### I don't have Docker

If you don't have a host to run docker, try [Koyeb](koyeb.com) to cloud-host your relays container.  The first container is free (account required).  Click the deploy button to get started.

[![Deploy to Koyeb](https://www.koyeb.com/static/images/deploy/button.svg)](https://app.koyeb.com/deploy?type=docker&image=ghcr.io/merliot/relays:main&name=relays&instance_type=free&ports=8000;http;/&env[WS_SCHEME]=wss://)
