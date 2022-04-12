package model

type Example struct {
	// gorm.Model
	ID       int
	Path     string
	NodeName string
}

func (Example) TableName() string {
	return "example"
}
