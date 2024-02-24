package main

import (
	"flag"
	"fmt"
)

func main() {
	portPtr := flag.Int("port", 3000, "Default port")
	flag.Parse()
	fmt.Println("my application listen on:", *portPtr)
}
