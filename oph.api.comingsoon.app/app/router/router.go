package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"oph.api.comingsoon/oph.api.comingsoon.app/app/controller/healthcheckctrl"
	"oph.api.comingsoon/oph.api.comingsoon.app/app/mdlwr"
	infrastructure "oph.api.comingsoon/oph.api.comingsoon.infrastructure"
	ioc "oph.api.comingsoon/oph.api.comingsoon.ioc"
)

type Router struct {
	infra      infrastructure.IInfrastructure
	ioc        ioc.IIoc
	middleware mdlwr.MiddlewareAPI
}

type IRouter interface {
	BuildRouters(app *echo.Echo)
}

func NewRouter(infra infrastructure.IInfrastructure, ioc ioc.IIoc) Router {
	return Router{
		infra,
		ioc,
		mdlwr.New(infra.Config(), infra.Logger()),
	}
}

func (r Router) BuildRouters(app *echo.Echo) {
	echo.New()
	r.buildHomeRouter(app)
	r.buildSwaggerRouters(app)
	r.buildHealthCheckRouters(app, app.Group("/HealthCheck"))
}

func (r Router) buildHomeRouter(app *echo.Echo) {
	app.GET("/", func(c echo.Context) error {
		return c.String(200, "\n\n\n\n\nnew\nproject\ncoming\nsoon\n\n\n\n\ndeveloped by onlineproducthouse.com")
	})
}

func (r Router) buildSwaggerRouters(app *echo.Echo) {
	if r.infra.Config().RunSwagger() {
		r.infra.Logger().Info("swagger enabled")
		app.GET("/swagger/*", echoSwagger.WrapHandler, r.middleware.NotImplemented("production"))
	}
}

func (r Router) buildHealthCheckRouters(app *echo.Echo, apiGroupV1 *echo.Group) {
	ctrl := healthcheckctrl.New(r.infra.Logger(), r.ioc.HealthCheck(), r.infra.Config())
	apiGroupV1.GET("/Ping", ctrl.Ping)
}
