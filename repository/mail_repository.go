package repository

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/erdogancayir/nargileapi/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mail.v2"
)

type mailRepository struct {
	database   *mongo.Database
	collection string
}

// Bu userRepository yapısı,
//domain.UserRepository arayüzünde tanımlanan tüm metodları uyguladığı için,
//bu tür bir değeri döndürmek mümkün.
func NewMailRepository(db *mongo.Database, collection string) domain.MailRepository {
	return &mailRepository{
		database:   db,
		collection: collection,
	}
}

func (a *mailRepository) SendMail(to string) (string, error) {
	// Doğrulama kodu oluşturma
	verificationCode := fmt.Sprintf("%06d", rand.Intn(1000000))
	msg := "From: no-reply\n" +
		"To: " + to + "\n" +
		"Subject: Verification Code\n\n" +
		"Your verification code is: " + verificationCode

	m := mail.NewMessage()
	m.SetHeader("From", "e.cayir2022@gtu.edu.tr")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Confirmation code")
	m.SetBody("text/plain", msg)

	d := mail.NewDialer("smtp.office365.com", 587, "e.cayir2022@gtu.edu.tr", "boston1907..")
	d.StartTLSPolicy = mail.MandatoryStartTLS

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
		return "", err
	}
	return verificationCode, nil
}
