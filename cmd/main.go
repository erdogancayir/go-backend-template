package main

import (
	"log"
	"time"

	"github.com/erdogancayir/nargileapi/api/route"
	"github.com/erdogancayir/nargileapi/bootstrap"
	"github.com/labstack/echo/v4"
)

func main() {

	// Database bağlantısı oluşturulur.
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	if db == nil {
		log.Println("Cannot created Connecting to DB")
		return
	}
	defer app.CloseDBConnection()
	timeout := time.Duration(env.ContextTimeout) * time.Second
	e := echo.New()
	/* e.Use(middleware.Logger())
	e.Use(middleware.Recover()) */
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			origin := req.Header.Get(echo.HeaderOrigin)
			res.Header().Set(echo.HeaderAccessControlAllowOrigin, origin)
			res.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
			res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			res.Header().Set("Access-Control-Allow-Credentials", "true")
			return next(c)
		}
	})

	route.Setup(env, timeout, db, e)

	e.Start(env.ServerAddress)
}
