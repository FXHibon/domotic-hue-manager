package app

import (
	"github.com/FXHibon/domotic-hue-manager/client"
	"github.com/FXHibon/domotic-hue-manager/config"
	"log"
)

func Run() {

	conf := config.LoadConfigurationOrDie()

	log.Printf("Will be using %#v\n", conf)

	err := client.ValidateCredentials(conf.Api)

	if err != nil {
		panic(err)
	}

}
