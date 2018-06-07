package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"strconv"
	"sync"

	color "github.com/fatih/color"
	telnet "github.com/reiver/go-telnet"
)

var domains = [...]string{"sunnyctl.povedaingenieria.com", "api.sunnyctl.io", "cloud.sunnyctl.io"}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go tunnel(&wg)
	go broker(&wg)
	go lookup(&wg)
	wg.Wait()
}

func tunnel(waitGrp *sync.WaitGroup) {
	color.Magenta("Conexion to tunnel server ")

	conn, err := telnet.DialTo("54.207.26.75:22")
	if nil != err {
		str := html.UnescapeString("&#" + strconv.Itoa(128683) + ";")
		fmt.Println(str)
		fmt.Println(err)
	} else {
		conn.Close()
		str := html.UnescapeString("&#" + strconv.Itoa(9989) + ";")
		fmt.Println(str)

	}
	defer waitGrp.Done()
}

func broker(waitGrp *sync.WaitGroup) {
	color.Magenta("Conexion to MQTT broker ")

	conn, err := telnet.DialTo("sunnyctl.povedaingenieria.com:1883")
	if nil != err {
		str := html.UnescapeString("&#" + strconv.Itoa(128683) + ";")
		fmt.Println(str)
	} else {
		conn.Close()
		str := html.UnescapeString("&#" + strconv.Itoa(9989) + ";")
		fmt.Println(str)

	}
	defer waitGrp.Done()

}

func lookup(waitGrp *sync.WaitGroup) {
	color.Magenta("DNS resolution tests ")
	for _, element := range domains {
		log.Println(net.LookupHost(element))
	}
	defer waitGrp.Done()

}
