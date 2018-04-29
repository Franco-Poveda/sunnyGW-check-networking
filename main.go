package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"strconv"

	rainbow "github.com/raphamorim/go-rainbow"
	telnet "github.com/reiver/go-telnet"
)

var domains = [...]string{"sunnyctl.povedaingenieria.com", "api.sunnyctl.io", "cloud.sunnyctl.io"}

func main() {

	tunnel()
	broker()
	lookup()
}

func tunnel() {
	fmt.Printf("%s", rainbow.Bold(rainbow.Magenta("Conexion to tunnel server ")))

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
}

func broker() {
	fmt.Printf("%s", rainbow.Bold(rainbow.Magenta("Conexion to MQTT broker ")))

	conn, err := telnet.DialTo("sunnyctl.povedaingenieria.com:1883")
	if nil != err {
		str := html.UnescapeString("&#" + strconv.Itoa(128683) + ";")
		fmt.Println(str)
	} else {
		conn.Close()
		str := html.UnescapeString("&#" + strconv.Itoa(9989) + ";")
		fmt.Println(str)

	}
}

func lookup() {
	fmt.Printf("%s", rainbow.Bold(rainbow.Magenta("DNS resolution tests ")))
	for _, element := range domains {
		log.Println(net.LookupHost(element))
	}
}
