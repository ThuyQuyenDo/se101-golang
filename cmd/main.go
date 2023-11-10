package main

import (
	"flag"
	"go-ecommerce/cmd/app"
)

func main() {
	app.Run(&app.FlagArg{
		ConfigPath: flag.String("config", "", "application config path"),
		Env:        flag.String("env", "", "application env"),
		Upgrade:    flag.Bool("upgrade", false, "ipgrade database"),
	})
}

// Command:  go run cmd/main.go -config ./config/env -env=local -upgrade=false
