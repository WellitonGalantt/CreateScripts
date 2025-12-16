package script

type ScriptDTOInput struct {
	UserId          int    `json:"user_id"`
	Type            string `json:"type"`
	Style           string `json:"style"`
	Topic           string `json:"topic"`
	Content         string `json:"content"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Hashtags        string `json:"hashtags"`
	ThumbnailPrompt string `json:"thumbnail_prompt"`
	IsFavorite      bool   `json:"is_favorite"`
}
