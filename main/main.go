package main

import (
	"fmt"
	"main/router"
	"udp_info"
)

const kHost = "0.0.0.0"
const kPort = 8888

func main() {
	go udp_info.Start()
	r := router.Start()
	err := r.Run(fmt.Sprintf("%s:%d", kHost, kPort))
	if err != nil {
		panic(err)
	}
}
