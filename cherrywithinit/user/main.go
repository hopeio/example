package main

import (
	"github.com/hopeio/cherry"
	"github.com/hopeio/cherry/_example/user/api"
	"github.conm/hopeio/collection/cherrywithinit/confdao"

	"github.com/hopeio/initialize"
	"github.com/hopeio/initialize/conf_center/nacos"
)

func main() {
	defer initialize.Start(confdao.Conf, confdao.Dao, nacos.ConfigCenter)()
	confdao.Conf.Server.WithOptions(cherry.WithGrpcHandler(api.GrpcRegister), cherry.WithGinHandler(api.GinRegister)).Run()
}
