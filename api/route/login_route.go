package route

import (
	"time"

	"github.com/erdogancayir/nargileapi/api/controller"
	"github.com/erdogancayir/nargileapi/bootstrap"
	"github.com/erdogancayir/nargileapi/domain"
	"github.com/erdogancayir/nargileapi/repository"
	"github.com/erdogancayir/nargileapi/usecase"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, pubGroup *echo.Group) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	pubGroup.POST("/login", lc.Login)
}
