package models

type Tracked struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Category    string `json:"category"`
	Genre       string `json:"genre"`
	Rating      string `json:"rating"`
	UserID      int64  `json:"userID"`
}
