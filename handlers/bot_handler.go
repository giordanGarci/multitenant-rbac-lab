package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/giordanGarci/api-tenants/interceptors"
	"github.com/giordanGarci/api-tenants/services"
	"github.com/giordanGarci/api-tenants/structs"
)

type BotHandler struct {
	service *services.Service
}

// NewBotHandler recebe o service já configurado (Dependency Injection)
func NewBotHandler(service *services.Service) *BotHandler {
	return &BotHandler{
		service: service,
	}
}

// GetAllBotsHandler godoc
// @Summary      Listar todos os robôs
// @Description  Retorna uma lista de todos os robôs cadastrados
// @Tags         bots
// @Produce      json
// @Success      200  {array}  structs.Bot
// @Router       /bots [get]
func (h *BotHandler) GetAllBotsHandler(w http.ResponseWriter, r *http.Request) {

	user, ok := interceptors.FromContext(r.Context())
	if !ok {
		http.Error(w, "user not found in context", http.StatusUnauthorized)
		return
	}

	tenantID, ok := interceptors.TenantFromContext(r.Context())
	if !ok {
		http.Error(w, "tenant not found in context", http.StatusUnauthorized)
		return
	}

	fmt.Println("User ID from context:", user.UserID)
	fmt.Println("Tenant ID from context:", tenantID)

	bots, _ := h.service.GetBots(tenantID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bots)
}

// GetBotByIDHandler godoc
// @Summary      Obter robô por ID
// @Description  Retorna os detalhes de um robô específico pelo seu ID
// @Tags         bots
// @Produce      json
// @Param        id   query      int  true  "ID do Robô"
// @Success      200  {object}  structs.Bot
// @Router       /bot [get]
func (h *BotHandler) GetBotByIDHandler(w http.ResponseWriter, r *http.Request) {
	botID := r.URL.Query().Get("id")
	botIDInt, err := strconv.ParseInt(botID, 10, 64)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	bot, _ := h.service.GetBot(botIDInt)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bot)
	w.WriteHeader(http.StatusOK)
}

// CreateBotHandler godoc
// @Summary      Criar um novo robô
// @Description  Cria um robô passando nome e status
// @Tags         bots
// @Accept       json
// @Produce      json
// @Param        robot  body      structs.Bot  true  "Dados do Robô"
// @Success      201  {string}  string "Created"
// @Router       /bot/create [post]
func (h *BotHandler) CreateBotHandler(w http.ResponseWriter, r *http.Request) {
	var bot structs.Bot
	if err := json.NewDecoder(r.Body).Decode(&bot); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.service.CreateBot(bot); err != nil {
		http.Error(w, "could not create bot", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
