package components

var component_listener_id uint = 0

type ComponentListener struct {
	Id       uint                      `json:"id"`
	Once     bool                      `json:"once"`
	Callback ComponentListenerCallback `json:"callback"`
}

type ComponentListenerCallback func()

func ToComponentListener(callback ComponentListenerCallback, once bool) ComponentListener {

	var listener ComponentListener

	listener.Id       = component_listener_id
	listener.Once     = once
	listener.Callback = callback

	component_listener_id += 1

	return listener

}
