package main

import (
	"testproject/util"

	"github.com/sirupsen/logrus"
)

func main() {
	util.GetData()
	logrus.Infof("Data received...yay!")
}
