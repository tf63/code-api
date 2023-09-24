package entity

type Todo struct {
	Id   uint
	Title string
	Done bool
}

type NewTodo struct {
	Title string
}
