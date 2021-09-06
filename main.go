package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/peterhellberg/ruuvitag"
)

var appConfig conf

func main() {
	log.Println("starting ruuvigw-go")
	appConfig.getConf()
	connectToInflux()
	go setupPoller()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func setupPoller() {
	for range time.Tick(time.Second * time.Duration(appConfig.PollInterval)) {
		for _, ip := range appConfig.GatewayIPs {
			go poll(ip)
		}
	}
}

type gwResponse struct {
	Data gwData
}
type gwData struct {
	Coordinates string
	Timestamp   int
	Gw_mac      string
	Tags        map[string]gwTagData
}
type gwTagData struct {
	Rssi      int
	Timestamp int
	Data      string
}

func poll(ip string) {
	resp, err := http.Get(fmt.Sprintf("http://%s/history?time=%d", ip, appConfig.PollInterval))
	if err != nil {
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	sb := string(body)
	var gwResp gwResponse
	json.Unmarshal([]byte(sb), &gwResp)
	for k := range gwResp.Data.Tags {
		//log.Println(k)
		data, err := hex.DecodeString(gwResp.Data.Tags[k].Data)
		if err != nil {
			continue
		}
		raw, err := ruuvitag.ParseRAWv2(data[5:])
		if err != nil {
			raw1, err := ruuvitag.ParseRAWv1(data[5:])
			if err != nil {
				continue
			}
			//log.Printf("%+v\n", raw1)
			insertRaw1(ip, gwResp.Data.Gw_mac, k, gwResp.Data.Tags[k].Rssi, raw1)
		} else {
			//log.Printf("%+v\n", raw)
			insertRaw2(ip, gwResp.Data.Gw_mac, k, gwResp.Data.Tags[k].Rssi, raw)
		}
	}
}
