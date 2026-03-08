package dto

type PasteRequest struct {
	Content    string `json:"content" binding:"required"`
	Visibility string `json:"visibility"`
	Expiry     string `json:"exp"`
}
