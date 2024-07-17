package main

import (
	"fmt"
	"github.com/hopeio/initialize"
	"github.com/hopeio/initialize/conf_center/nacos"
	"github.com/hopeio/initialize/conf_dao/mqtt"
	"time"
)

type Config struct {
	A    string `flag:"name:aaa;short:a;default:a;usage:模块名;env:AAA"`
	B    int
	T    time.Duration
	Auth Auth
}

func (c *Config) InitBeforeInject() {
	c.A = "A"
}

func (c *Config) InitAfterInject() {
	if c.A == "B" {
		c.A = "A"
	}
	if c.T < time.Second {
		c.T = c.T * time.Second
	}
}

type Auth struct {
	Token string
}

type Dao struct {
	Mqtt mqtt.Client
	//DB   postgres.DB
}

func (d *Dao) InitBeforeInject() {

}

func (d *Dao) InitAfterInjectConfig() {

}

func (d *Dao) InitAfterInject() {
	if token := d.Mqtt.Publish("test", 0, false, "test"); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

var config Config
var dao Dao

func main() {
	defer initialize.Start(&config, &dao, &nacos.Nacos{})()
	fmt.Println(config)
	initialize.RegisterDeferFunc(func() {
		fmt.Println("defer")
	})
}
