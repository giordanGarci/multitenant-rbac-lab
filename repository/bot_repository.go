package repository

import (
	"sync"

	"github.com/giordanGarci/api-tenants/structs"
)

// BotRepository define o contrato para acesso a dados de bots
// Qualquer implementação (memória, banco, API) deve seguir este contrato
type BotRepository interface {
	GetAllBots(tenantID int64) ([]structs.Bot, error)
	GetBotByID(id int64) (*structs.Bot, error)
	AddBot(bot structs.Bot) error
}

// InMemoryBotRepository é a implementação em memória do BotRepository
type InMemoryBotRepository struct {
	bots []structs.Bot
	mu   sync.Mutex
}

// NewInMemoryBotRepository cria uma nova instância do repositório em memória
func NewInMemoryBotRepository() BotRepository {
	return &InMemoryBotRepository{
		bots: []structs.Bot{
			{ID: 1, Name: "BotOne", Status: "active", TenantId: 1001},
			{ID: 2, Name: "BotTwo", Status: "inactive", TenantId: 1002},
			{ID: 3, Name: "BotThree", Status: "active", TenantId: 1003},
		},
	}
}

func (r *InMemoryBotRepository) GetAllBots(tenantID int64) ([]structs.Bot, error) {
	var filteredBots []structs.Bot

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, bot := range r.bots {
		if bot.TenantId == tenantID {
			filteredBots = append(filteredBots, bot)
		}
	}
	return filteredBots, nil
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
