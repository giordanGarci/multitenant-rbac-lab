package services

import (
	"github.com/giordanGarci/api-tenants/repository"
	"github.com/giordanGarci/api-tenants/structs"
)

type Service struct {
	repository repository.BotRepository // Depende da INTERFACE, não da implementação
}

// NewService recebe a dependência por parâmetro (Dependency Injection)
func NewService(repo repository.BotRepository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) GetBots(tenantID int64) ([]structs.Bot, error) {

	bots, _ := s.repository.GetAllBots(tenantID)

	return bots, nil
}

func (s *Service) GetBot(id int64) (*structs.Bot, error) {
	bot, _ := s.repository.GetBotByID(id)
	return bot, nil
}

func (s *Service) CreateBot(bot structs.Bot) error {
	return s.repository.AddBot(bot)
}
