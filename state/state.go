package state

// Transformation perform a read to the application state
type Transformation interface {
	Run(ApplicationState) ApplicationState
}

// AppendFilterString adds to a filter string
type AppendFilterString struct {
	Value string
}

// ChopFilterString adds to a filter string
type ChopFilterString struct{}

func (t *AppendFilterString) Run(as ApplicationState) ApplicationState {
	as.FilterString = as.FilterString + t.Value
	return as
}

func (t *ChopFilterString) Run(as ApplicationState) ApplicationState {
	sz := len(as.FilterString)
	as.FilterString = as.FilterString[:sz-1]
	return as
}

// ApplicationState the state of the application
type ApplicationState struct {
	FilterString string
}

// Application contains the structure of the application data, including state and read and write queues.
type Application struct {
	State      ApplicationState
	WriteQueue []Transformation
}

func (a *Application) write(t Transformation) {
	a.WriteQueue = append(a.WriteQueue, t)
}

func (a *Application) read() ApplicationState {
	return a.State
}

func createApplicationState() Application {
	a := Application{
		State:      ApplicationState{FilterString: ""},
		WriteQueue: make([]Transformation, 0),
	}

	return a
}
