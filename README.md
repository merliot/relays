# Relays Device

Relays device controls 1-4 relays using a microcontroller or SBC, such as Rapsberry Pi.

<div align="center">
  
![](images/nano-rp2040-relays.png "Arduino Nano rp2040 Connect with 4 relays")

</div>

## Demo

Try the live [demo](https://demo.merliot.net/relays-01/).  Click the &#x1F4E5; button to see how to build and deploy a relays device.

<div align="center">

[![](images/relays-01.png)](https://demo.merliot.net/relays-01/)

</div>

## Running on Merliot Hub

A [Merliot Hub](https://github.com/merliot/hub) is the easiest way to run the relays device.

## Run Standalone with Docker

```
git clone https://github.com/merliot/relays.git
cd relays
docker build -t relays .
docker run -p 8000:8000 relays
```

Now the device is now running in a docker container as a web server, listening on port :8000.  

Browse to http://\<host\>:8000 to view and setup the device.

If the docker host is using https://, pass in the environment variable WS_SCHEME=wss://.

```
docker run -e "WS_SCHEME=wss://" -p 8000:8000 relays
```

See full list of [environment](https://github.com/merliot/device/blob/main/docs/environment.md) variables.



## Run from Source Code

```
git clone https://github.com/merliot/relays.git
cd relays
PORT_PRIME=8000 go run -tags prime ./cmd
```

Browse to http://\<host\>:8000 to view and deploy device.
