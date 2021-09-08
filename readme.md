This is a small program to get data from ruuvi gateways to influx.

For this to work you need

- Influxdb 2.x
- A RuuviGateway configured with no autentication


## How to use:

### Start influx2
`docker run --name ruuvi-influx -d -p 8086:8086 influxdb`

### Go to influx config run the setup http://localhost:8086
### Edit the conf.yaml with your influx token/org/bucket & gateway IP(s)
ie:
```
influxHost: http://localhost:8086
influxToken: qI9RqpoUTpMS6-OXpmQZjKXwEsbuaZV3Mw3anNpM-RBmOa__nT8hAhFnkqi5NbzOCMJ45LKbsuWbG-uH0ViPUQ== 
influxBucket: ruuviBucket
influxOrg: ruuviOrg
pollInterval: 2
gatewayIPs:
  - "10.99.20.159"
```

### Build ruuvigw-go
`docker build --tag ruuvigw-go .`

## Run it!

`docker run --name ruuvigw-go --restart unless-stopped -d --net=host ruuvigw-go:latest`
