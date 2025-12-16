package script

import "time"

type Script struct {
	ID              int       `json:"id"`
	UserId          int       `json:"user_id"`
	Type            string    `json:"type"`
	Style           string    `json:"style"`
	Topic           string    `json:"topic"`
	Content         string    `json:"content"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Hashtags        string    `json:"hashtags"`
	ThumbnailPrompt string    `json:"thumbnail_prompt"`
	IsFavorite      bool      `json:"is_favorite"`
	CreatedAt       time.Time `json:"created_at"`
}
