package stack

type stack[T any] struct {
	data []T
}

type Stack[T any] interface {
	Push(data T)
	Pop() T
	Get() []T
}

func New[T any]() Stack[T] {
	return &stack[T]{
		data: []T{},
	}
}

func (s *stack[T]) Push(data T) {
	s.data = append(s.data, data)
}

func (s *stack[T]) Pop() T {
	lastData := s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]
	return lastData
}

func (s *stack[T]) Get() []T {
	return s.data
}
