package route

import (
	"time"

	"github.com/erdogancayir/nargileapi/bootstrap"
	"github.com/erdogancayir/nargileapi/middleware"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, e *echo.Echo) {
	//ortak middleware'leri belirlemek için
	publicRouter := e.Group("")
	NewSignupRouter(env, timeout, db, publicRouter)
	NewMailRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := e.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
}
