package main

import (
	"fmt"
	"html"
	"strconv"

	"github.com/raphamorim/go-rainbow"
	"github.com/reiver/go-telnet"
)

func main() {
	fmt.Printf("%s", rainbow.Bold(rainbow.Magenta("Conexion to tunnel server ")))

	conn, err := telnet.DialTo("54.207.26.75:22")
	if nil != err {
		str := html.UnescapeString("&#" + strconv.Itoa(128683) + ";")
		fmt.Println(str)
	} else {
		conn.Close()
		str := html.UnescapeString("&#" + strconv.Itoa(9989) + ";")
		fmt.Println(str)

	}
}