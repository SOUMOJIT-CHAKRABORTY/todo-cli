package todo

import (
	"errors"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {

	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t

	if index < 0 || index >= len(ls) {
		return errors.New("Invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t

	if index < 0 || index >= len(ls) {
		return errors.New(("Invalid index"))
	}

	// First argument to append must be original slice; have []item {form 0th index to index-2, excluding index-1}
	// Second argument to append must be new elements to be added; {form index to last index}
	*t = append(ls[:index-1], ls[index:]...)

	return nil
}
