package done

import "github.com/LakshyaNegi/todos/internal/entity"

func NewDoneModelFromTodos(todos []*entity.Todo) model {
	return model{
		choices:  todos,
		selected: make(map[int]*entity.Todo),
	}
}
