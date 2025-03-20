package components

import gooey "github.com/cookiengineer/gooey/pkg"
import "github.com/cookiengineer/gooey/pkg/dom"
import "strings"

type Label struct {
	Label string `json:"label"`
	Type  string `json:"type"`
	Component
}

func NewLabel(label string, typ string) Label {

	var component Label

	element := gooey.Document.CreateElement("label")

	component.Label = label
	component.Type  = strings.ToLower(typ)

	component.Init(element)
	component.Render()

	return component

}

func ToLabel(element *dom.Element) Label {

	var component Label

	component.Label = strings.TrimSpace(element.TextContent)
	component.Type  = element.GetAttribute("data-type")

	component.Init(element)

	return component

}

func (component *Label) Render() {

	if component.Element != nil {

		if component.Label != "" {
			component.Element.SetInnerHTML(component.Label)
		}

		if component.Type != "" {
			component.Element.SetAttribute("data-type", component.Type)
		}

	}

}

func (component *Label) String() string {

	html := ""

	if component.Type != "" {
		html += "<label data-type=\"" + component.Type + "\">"
	} else {
		html += "<label>"
	}

	if component.Label != "" {
		html += component.Label
	}

	html += "</label>"

	return html

}
