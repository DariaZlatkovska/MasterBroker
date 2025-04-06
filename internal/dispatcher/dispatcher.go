package dispatcher

type Event interface {
	EventType() string
}

type Handler interface {
	Handle(event Event) error
}

type Dispatcher struct {
	handlers map[string]Handler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{handlers: make(map[string]Handler)}
}

func (d *Dispatcher) Register(eventType string, handler Handler) {
	d.handlers[eventType] = handler
}

func (d *Dispatcher) Dispatch(event Event) error {
	handler, ok := d.handlers[event.EventType()]
	if !ok {
		return nil
	}
	return handler.Handle(event)
}
