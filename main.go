package main

import (
	"testproject/util"

	log "github.com/sirupsen/logrus"
)

func main() {
	util.GetData()
	log.Infof("Data received...yay!")
}
