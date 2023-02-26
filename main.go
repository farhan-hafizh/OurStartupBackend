package main

import (
	"flag"
	"ourstartup/server"
)

func main() {
	mode := flag.String("mode", "development", "For environtment variables")
	flag.Parse()

	server.Init(*mode)
}
