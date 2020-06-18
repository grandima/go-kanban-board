package entity

type Task struct {
	Id int 		`json:"id"`
	Name string	 `json:"name" validate:"max=500,min=1"`
	Description string `json:"description" validate:"max=5000"`
}