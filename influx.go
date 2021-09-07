package main

import (
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"github.com/influxdata/influxdb-client-go/api/write"
	"github.com/peterhellberg/ruuvitag"
)

var dbClient influxdb2.Client
var writeAPI api.WriteAPI

func connectToInflux() {
	dbClient = influxdb2.NewClient(appConfig.InfluxHost, appConfig.InfluxToken)
	writeAPI = dbClient.WriteAPI(appConfig.InfluxOrg, appConfig.InfluxBucket)
}

func writePoint(point *write.Point) {
	writeAPI.WritePoint(point)
}

func insertRaw2(gwip string, gwmac string, mac string, rssi int, data ruuvitag.RAWv2) {
	writePoint(influxdb2.NewPointWithMeasurement("stat").
		AddTag("gateway-ip", gwip).
		AddTag("gateway-mac", gwmac).
		AddTag("mac", mac).
		AddField("temperature", data.Temperature).
		AddField("humidity", data.Humidity).
		AddField("pressure", data.Pressure).
		AddField("movementCounter", data.Movement).
		AddField("dataFormat", data.DataFormat).
		AddField("battery", data.Battery).
		AddField("accelerationX", data.Acceleration.X).
		AddField("accelerationY", data.Acceleration.Y).
		AddField("accelerationZ", data.Acceleration.Z).
		AddField("measurementSequenceNumber", data.Sequence).
		AddField("txPower", data.TXPower).
		AddField("rssi", rssi))
}

func insertRaw1(gwip string, gwmac string, mac string, rssi int, data ruuvitag.RAWv1) {
	writePoint(influxdb2.NewPointWithMeasurement("stat").
		AddTag("gateway-ip", gwip).
		AddTag("gateway-mac", gwmac).
		AddTag("mac", mac).
		AddField("temperature", data.Temperature).
		AddField("humidity", data.Humidity).
		AddField("pressure", data.Pressure).
		AddField("dataFormat", data.DataFormat).
		AddField("battery", data.Battery).
		AddField("accelerationX", data.Acceleration.X).
		AddField("accelerationY", data.Acceleration.Y).
		AddField("accelerationZ", data.Acceleration.Z).
		AddField("rssi", rssi))
}
