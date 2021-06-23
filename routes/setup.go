package routes

import (
	"fmt"
	"full-stack-go/app"
	"full-stack-go/constants"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	port string
}

func (r *Router) SetupRouter(env *app.Environment) {
	r.port = env.GetAPIPort()

	switch env.EnvMode {
	case constants.DEV:
		gin.SetMode(gin.DebugMode)
		r.engine = gin.Default()
		DevRoutes(r)
	case constants.TEST:
		gin.SetMode(gin.TestMode)
		r.engine = gin.Default()
		TestRoutes(r)
	case constants.PRODUCTION:
		gin.SetMode(gin.ReleaseMode)
		r.engine = gin.Default()
		ProductionRoutes(r)
	default:
		panic(fmt.Errorf("environment mode: %s invalid. The environments are required: [dev, test, production]", env.EnvMode))
	}
}

func (r *Router) Run() {
	//run on all interfaces
	r.engine.Run(fmt.Sprintf(":%s", r.port))
}