package main

import (
	"CRUD_SERVER/init/cmd"
	"flag"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
