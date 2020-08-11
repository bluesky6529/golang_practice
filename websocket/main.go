package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

func main() {
	b, err := ioutil.ReadFile("addr.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, addr := range strings.Split(strings.ReplaceAll(string(b), "\r\n", "\n"), "\n") {
		conn, res, err := websocket.DefaultDialer.Dial(addr, nil)
		if err != nil {
			color.Red("connect to %s, err: %s", addr, err)
		} else {
			color.Green("connect to %s, res: %+v", addr, res)
			conn.Close()
		}
	}
	color.Cyan("Press the any key to terminate.")
	var input string
	fmt.Scanln(&input)
}
