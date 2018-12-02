package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/denislee/stack-backend/config"
	"github.com/denislee/stack-backend/server"
)

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)
	server.Init()
}
