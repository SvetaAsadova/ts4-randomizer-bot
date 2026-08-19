package db

// Event представляет событие из базы данных.
type Event struct {
	ID          int64  `json:"id"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Variables   string `json:"variables"`
	DLC         sqlNullString
	Impact      int    `json:"impact"`
	CreatedAt   string `json:"created_at"`
}

// sqlNullString — удобная обёртка для NULL-текстовых полей.
type sqlNullString struct {
	Value string
	Valid bool
}

// FormatAsCard возвращает строку для Telegram с форматированием карточки:
// **заголовок**
//
// описание
func (e Event) FormatAsCard() string {
	result := "**" + e.Title + "**"
	if e.Description != "" {
		result += "\n\n" + e.Description
	}
	return result
}
