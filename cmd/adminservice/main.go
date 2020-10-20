package main

import (
	"flag"
)

var configFile = flag.String("c", "configs/adminservice.yaml", "set config file which viper will loading.")

func main() {

	flag.Parse()

	app, err := CreateApps(*configFile)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()

}
