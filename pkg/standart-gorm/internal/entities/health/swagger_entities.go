package health

type ResponseHealth struct {
	Status int    `json:"status" example:"200"`
	Result string `json:"result" example:"Service OK"`
}
