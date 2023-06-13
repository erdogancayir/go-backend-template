package usecase

import (
	"time"

	"github.com/erdogancayir/nargileapi/domain"
)

type mailUsecase struct {
	mailRepository domain.MailRepository
	contextTimeout time.Duration
}

func NewMailUsecase(mailRepository domain.MailRepository, timeout time.Duration) domain.MailRepository {
	return &mailUsecase{
		mailRepository: mailRepository,
		contextTimeout: timeout,
	}
}

func (mu *mailUsecase) SendMail(to string) (string, error) {
	return mu.mailRepository.SendMail(to)
}
