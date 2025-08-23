package main

import (
	"github.com/NoANameGroup/The-DAOld-Backend/adaptor/router"
	"github.com/NoANameGroup/The-DAOld-Backend/infra/util/log"
	"github.com/NoANameGroup/The-DAOld-Backend/provider"
)

func Init() {
	provider.Init()
	log.Info("所有模块初始化完成...")
}

func main() {
	Init()
	r := router.SetupRoutes()

	log.Info("服务器即将启动于 :8080")
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
