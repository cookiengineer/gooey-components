package app

type Component struct {
	listeners map[string][]*ComponentListener `json:"listeners"`
}

func (component *Component) InitEvent(event string) {

	_, ok := component.listeners[event]

	if ok == false {
		component.listeners[event] = make([]*Listener, 0)
	}

}

func (component *Component) AddEventListener(event string, listener ComponentListener) bool {

	var result bool = false

	_, ok := component.listeners[event]

	if ok == true {
		component.listeners[event] = append(component.listeners[event], &listener)
		result = true
	}

	return result

}

func (component *Component) FireEventListeners(event string) bool {

	var result bool = false

	listeners, ok := component.listeners[event]

	if ok == true {

		indexes := make([]int, 0)

		for l := 0; l < len(listeners); l++ {

			listener := listeners[l]

			listener.Callback()

			if listener.Once == true {
				indexes = append(indexes, l)
			}

		}

		if len(indexes) > 0 {

			for _, index := range indexes {
				listeners = append(listeners[:index], listeners[index+1:]...)
			}

			component.listeners[event] = listeners

		}

	}

	return result

}

func (component *Component) RemoveEventListener(event string, listener *ComponentListener) bool {

	var result bool = false

	if listener != nil {

		listeners, ok := component.listeners[event]

		if ok == true {

			var index int = -1

			for l := 0; l < len(listeners); l++ {

				if listeners[l].Id == listener.Id {
					index = l
					break
				}

			}

			if index != -1 {
				component.listeners[event] = append(component.listeners[event][:index], component.listeners[event][index+1:]...)
				result = true
			}

		}

	} else {

		_, ok := component.listeners[event]

		if ok == true {
			component.listeners[event] = make([]*ComponentListener, 0)
			result = true
		}

	}

	return result

}

func (component *Component) Render() string {
	return "<component></component>"
}
