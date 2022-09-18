package stateroutine

type Operation int64

const (
	GET    Operation = 0
	SET    Operation = 1
	DELETE Operation = 2
)

type Message[T any] struct {
	operation Operation
	key       string
	value     *T
	channel   *chan *T
}

type State[T any] struct {
	channel chan Message[T]
}

func loop[T any](hash map[string]T, state State[T]) {
	for {
		msg := <-state.channel
		switch {
		case msg.operation == GET:
			var value_ptr *T = nil
			value, ok := hash[msg.key]
			if ok {
				value_ptr = &value
			}
			*msg.channel <- value_ptr
		case msg.operation == SET:
			hash[msg.key] = *msg.value
		case msg.operation == DELETE:
			delete(hash, msg.key)
		}
	}
}

func Go[T any]() State[T] {
	hash := make(map[string]T)
	state := State[T]{channel: make(chan Message[T], 100)}
	go loop(hash, state)
	return state
}

func Get[T any](state State[T], key string) *T {
	channel := make(chan *T)
	state.channel <- Message[T]{
		operation: GET,
		key:       key,
		value:     nil,
		channel:   &channel,
	}
	return <-channel
}

func Set[T any](state State[T], key string, value T) {
	state.channel <- Message[T]{
		operation: SET,
		key:       key,
		value:     &value,
		channel:   nil,
	}
}

func Delete[T any](state State[T], key string) {
	state.channel <- Message[T]{
		operation: DELETE,
		key:       key,
		value:     nil,
		channel:   nil,
	}
}
