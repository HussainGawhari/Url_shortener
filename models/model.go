package models

type Request struct {
	Id        int    `json:"id"`
	LongUrl   string `json:"long_url" validate:"required"`
	ShortUrl  string `json:"short_url"`
	CreatedAt string `json:"created_at"`
	// 	Expir_at string `form:"expir at" json:"expir at"`
}

type Response struct {
	Status  int64     `json:"status"`
	Message string    `json:"message"`
	Data    []Request `json:"data,omitempty"`
}
