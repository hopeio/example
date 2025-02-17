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

func (c *Config) BeforeInject() {
	c.A = "A"
}

func (c *Config) AfterInject() {
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

func (d *Dao) BeforeInject() {

}

func (d *Dao) AfterInjectConfig() {

}

func (d *Dao) AfterInject() {
	if token := d.Mqtt.Publish("test", 0, false, "test"); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

var global = initialize.NewGlobal[*Config, *Dao](nacos.ConfigCenter)

func main() {
	defer global.Cleanup()
	fmt.Println(global.Config)
	global.Defer(func() {
		fmt.Println("defer")
	})
}
