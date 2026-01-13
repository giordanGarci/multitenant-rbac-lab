package repository

import "github.com/giordanGarci/api-tenants/structs"

// BotRepository define o contrato para acesso a dados de bots
// Qualquer implementação (memória, banco, API) deve seguir este contrato
type BotRepository interface {
	GetAllBots() ([]structs.Bot, error)
	GetBotByID(id int64) (*structs.Bot, error)
	AddBot(bot structs.Bot) error
}

// InMemoryBotRepository é a implementação em memória do BotRepository
type InMemoryBotRepository struct {
	bots []structs.Bot
}

// NewInMemoryBotRepository cria uma nova instância do repositório em memória
func NewInMemoryBotRepository() BotRepository {
	return &InMemoryBotRepository{
		bots: []structs.Bot{
			{ID: 1, Name: "BotOne", Status: "active"},
			{ID: 2, Name: "BotTwo", Status: "inactive"},
		},
	}
}

func (r *InMemoryBotRepository) GetAllBots() ([]structs.Bot, error) {
	return r.bots, nil
}

func (r *InMemoryBotRepository) GetBotByID(id int64) (*structs.Bot, error) {
	for _, bot := range r.bots {
		if bot.ID == id {
			return &bot, nil
		}
	}
	return nil, nil
}

func (r *InMemoryBotRepository) AddBot(bot structs.Bot) error {
	r.bots = append(r.bots, bot)
	return nil
}
