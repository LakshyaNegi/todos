package del

import "github.com/LakshyaNegi/todos/internal/entity"

func NewDeleteModelFromTodos(todos []*entity.Todo) model {
	return model{
		choices:  todos,
		selected: make(map[int]*entity.Todo),
	}
}
