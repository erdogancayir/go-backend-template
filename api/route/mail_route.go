package route

import (
	"time"

	"github.com/erdogancayir/nargileapi/api/controller"
	"github.com/erdogancayir/nargileapi/bootstrap"
	"github.com/erdogancayir/nargileapi/repository"
	"github.com/erdogancayir/nargileapi/usecase"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMailRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, pubGroup *echo.Group) {
	//parametre olarak mongo ve users verileri gitti.
	//this file determines how users are stored and handled
	ur := repository.NewMailRepository(db, "mails")
	sc := controller.MailController{
		MailUsecase: usecase.NewMailUsecase(ur, timeout),
		Env:         env,
		Db:          db,
	}
	pubGroup.POST("/sendmail", sc.SendMail)
}
