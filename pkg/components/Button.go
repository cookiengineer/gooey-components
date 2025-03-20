package components

import gooey "github.com/cookiengineer/gooey/pkg"
import "github.com/cookiengineer/gooey/pkg/dom"
import "strings"

type Button struct {
	Label  string `json:"label"`
	Action string `json:"action"`
	Component
}

func NewButton(label string, action string) Button {

	var component Button

	element := gooey.Document.CreateElement("button")

	component.Label  = label
	component.Action = strings.ToLower(action)

	component.Init(element)
	component.Render()

	return component

}

func ToButton(element *dom.Element) Button {

	var component Button

	component.Label  = strings.TrimSpace(element.TextContent)
	component.Action = strings.ToLower(element.GetAttribute("data-action"))

	component.Init(element)

	return component

}

func (component *Button) Render() {

	if component.Element != nil {

		if component.Label != "" {
			component.Element.SetInnerHTML(component.Label)
		}

		if component.Action != "" {
			component.Element.SetAttribute("data-action", component.Action)
		}

	}

}

func (component *Button) String() string {

	html := ""

	if component.Action != "" {
		html += "<button data-action=\"" + component.Action + "\">"
	} else {
		html += "<button>"
	}

	if component.Label != "" {
		html += component.Label
	}

	html += "</button>"

	return html

}
