package main

import (
	"fmt"
	"github.com/hopeio/cherry/initialize"
)

type Config struct {
	Level string
}

func (c Config) InitBeforeInject() {

}

func (c Config) InitAfterInject() {

}

func main() {
	var config Config
	defer initialize.Start(&config, nil)()
	fmt.Println(config)
}
