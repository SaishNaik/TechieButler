package router

import (
	"github.com/gin-gonic/gin"
	"techiebutler/controller"
)

type Router struct {
	Engine     *gin.Engine
	controller *controller.Controller
}

func NewRouter(controller2 *controller.Controller) *Router {
	return &Router{
		Engine:     gin.Default(),
		controller: controller2,
	}
}

func (r *Router) InitRoutes() {
	r.Engine.GET("/records", r.controller.GetEmployeeRecords)
}
