package swissknife

type ChannelWorker struct {
	MaxCapacity          int
	Events               chan interface{}
	EventHandlerCallback func(event interface{})
	ProcessingEvents     bool
	IsFinished           bool

	isAsyncMode bool
}

func NewChannelWorker(callback func(event interface{}), maxCapacity int) *ChannelWorker {
	return &ChannelWorker{
		MaxCapacity:          maxCapacity,
		Events:               make(chan interface{}, maxCapacity),
		EventHandlerCallback: callback,
	}
}

func (w *ChannelWorker) SetAsync(asyncMode bool) *ChannelWorker {
	w.isAsyncMode = asyncMode
	return w
}

func (w *ChannelWorker) AddEvent(event interface{}) {
	// check channel capacity
	if len(w.Events) >= w.MaxCapacity {
		return
	}

	// add message to channel
	w.Events <- event
}

// Start handle events
// NOTE: it's blocking method
func (w *ChannelWorker) Start() {
	w.ProcessingEvents = true
	w.handleEvents()
}

func (w *ChannelWorker) Stop() {
	w.ProcessingEvents = false
}

func (w *ChannelWorker) GetMessagesAvailableCount() int {
	return len(w.Events)
}

func (w *ChannelWorker) handleEvents() {
	w.IsFinished = false
	for w.ProcessingEvents {
		for event := range w.Events {
			if !w.ProcessingEvents {
				break
			}

			if w.isAsyncMode {
				go w.EventHandlerCallback(event)
			} else {
				w.EventHandlerCallback(event)
			}
		}
	}
	w.ProcessingEvents = false
	w.IsFinished = true
}
