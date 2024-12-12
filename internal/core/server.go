package core

import (
	"gin-api/internal/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()
	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	address := ":8080"
	s := initServer(address, Router)
	//global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
