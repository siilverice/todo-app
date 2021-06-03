package Models

type Todo struct {
	ID          uint   `json:id`
	Title       string `json:"title" binding:"required"`
	Description string `json:description`
}

func (b *Todo) TableName() string {
	return "todo"
}