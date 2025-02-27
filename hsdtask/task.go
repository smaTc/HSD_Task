package hsdtask

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func NewTask(title, description string, completed bool) Task {
	return Task{Title: title, Description: description, Completed: completed}
}
