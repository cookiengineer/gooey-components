package app

var component_listener_id uint = 0

type ComponentListener struct {
	Id       uint                      `json:"id"`
	Callback ComponentListenerCallback `json:"callback"`
}

type ComponentListenerCallback func()

func ToComponentListener(callback ComponentListenerCallback) ComponentListener {

	var listener ComponentListener

	listener.Id = component_listener_id
	listener.Callback = callback

	component_listener_id += 1

	return listener

}
