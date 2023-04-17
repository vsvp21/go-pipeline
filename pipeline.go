package pipeline

import "errors"

// Stage defines interface for pipeline stage
type Stage[T any] interface {
	SetNext(s Stage[T])
	Next() Stage[T]
	Execute(obj T) error
}

// BaseStage stage with default behaviour
type BaseStage[T any] struct {
	next Stage[T]
}

func (s *BaseStage[T]) SetNext(n Stage[T]) {
	s.next = n
}

func (s *BaseStage[T]) Next() Stage[T] {
	return s.next
}

func (s *BaseStage[T]) Execute(obj T) error {
	panic("override")
}

// NewPipeline creates pipeline and collects stages in linked list
func NewPipeline[T any](stages []Stage[T]) (*Pipeline[T], error) {
	if len(stages) == 0 {
		return nil, errors.New("no stages provided")
	}

	head := stages[0]
	if len(stages) > 1 {
		current := head
		for i := 1; i < len(stages); i++ {
			current.SetNext(stages[i])
			current = stages[i]
		}
	}

	return &Pipeline[T]{Head: head}, nil
}

// Pipeline ...
type Pipeline[T any] struct {
	Head Stage[T]
}

func (p *Pipeline[T]) Execute(obj T) error {
	current := p.Head
	for current != nil {
		if err := current.Execute(obj); err != nil {
			return err
		}

		current = current.Next()
	}

	return nil
}
