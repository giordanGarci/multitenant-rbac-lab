package structs

type Bot struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	TenantId int64  `json:"tenant_id"`
}
