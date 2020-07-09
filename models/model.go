package models

// Todo your tasks that you have to
type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"string"`
	Description string `json:"description"`
}

// TableName just return todo
func (b *Todo) TableName() string {
	return "todo"
}
