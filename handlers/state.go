package handlers

type State struct {
	data   any
	state  string
	userId int64
}

type StateInterface interface {
	SetState(state string)
	GetState() string
	SetData(data any)
	GetData() any
}

func (st *State) SetState(state string) {
	st.state = state
}

func (st *State) GetState() string {
	return st.state
}

func (st *State) SetData(data any) {
	st.data = data
}

func (st *State) GetData() any {
	return st.data
}
