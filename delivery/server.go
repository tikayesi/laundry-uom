package delivery

import (
	"github.com/eulbyvan/go-enigma-laundry/config"
	"github.com/eulbyvan/go-enigma-laundry/delivery/controller"
	"github.com/eulbyvan/go-enigma-laundry/manager"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	engine         *gin.Engine
	useCaseManager manager.UseCaseManager
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	return &appServer{
		engine:         ginEngine,
		useCaseManager: usecase,
	}
}

func (a *appServer) initHandlers() {
	controller.NewUomController(a.engine, a.useCaseManager.UomUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(":8085")
	if err != nil {
		panic(err.Error())
	}
}
