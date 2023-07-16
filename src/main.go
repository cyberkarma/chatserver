package main

import (
	"github.com/cyberkarma/chatserver/cmd"
	"github.com/cyberkarma/chatserver/configs"
)

func init() {
	configs.AppConfig()
}

func main() {

	cmd.Execute()
}
