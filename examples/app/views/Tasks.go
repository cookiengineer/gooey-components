package views

import "example/actions"
import "example/schemas"
import gooey "github.com/cookiengineer/gooey/pkg"
import "github.com/cookiengineer/gooey/pkg/dom"
import "github.com/cookiengineer/gooey-components/pkg/app"
import "sort"
import "strconv"

type Tasks struct {
	Main   *app.Main      `json:"main"`
	Schema *schemas.Tasks `json:"schema"`
	app.BaseView
}

func NewTasks(main *app.Main) Tasks {

	var view Tasks

	view.Main = main
	view.Schema = &schemas.Tasks{}
	view.Elements = make(map[string]*dom.Element)

	view.SetElement("table",  gooey.Document.QuerySelector("main > table"))
	view.SetElement("dialog", gooey.Document.QuerySelector("dialog"))
	view.SetElement("header", gooey.Document.QuerySelector("header"))
	view.SetElement("footer", gooey.Document.QuerySelector("footer"))

	view.Init()

	return view

}

func (view Tasks) Init() {

	table  := view.GetElement("table")
	dialog := view.GetElement("dialog")
	footer := view.GetElement("footer")

	if table != nil {

		table.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

			target := event.Target

			if target.TagName == "INPUT" && target.GetAttribute("type") == "checkbox" {

				row      := target.QueryParent("tr")
				num, err := strconv.ParseInt(row.GetAttribute("data-id"), 10, 64)

				if err == nil {

					id       := int(num)
					task, ok := view.Schema.Tasks[id]

					if ok == true {

						if task.Done == true {
							task.Done = false
						} else {
							task.Done = true
						}

						go func() {

							actions.UpdateTask(view.Main.Client, task)
							view.Refresh()

						}()

					}

				}

			}

		}))

	}

	if dialog != nil {

		dialog.QuerySelector("button[data-action=\"close\"]").AddEventListener("click", dom.ToEventListener(func(event dom.Event) {
			view.CloseDialog()
		}))

		dialog.QuerySelector("button[data-action=\"cancel\"]").AddEventListener("click", dom.ToEventListener(func(event dom.Event) {
			view.CloseDialog()
		}))

		dialog.QuerySelector("button[data-action=\"confirm\"]").AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

			title := dialog.QuerySelector("input[data-name=\"title\"]").Value.Get("value").String()
			done  := dialog.QuerySelector("input[data-name=\"done\"]").Value.Get("checked").Bool()

			task := schemas.Task{
				ID: 0, // set by backend
				Title: title,
				Done:  done,
			}

			if task.Title != "" {

				buttons := dialog.QuerySelectorAll("button")

				for _, button := range buttons {
					button.SetAttribute("disabled", "")
				}

				go func() {

					actions.CreateTask(view.Main.Client, &task)
					view.CloseDialog()
					view.Refresh()

					for _, button := range buttons {
						button.RemoveAttribute("disabled")
					}

				}()

			}

		}))

	}

	if footer != nil {

		footer.QuerySelector("button[data-action=\"create\"]").AddEventListener("click", dom.ToEventListener(func(event dom.Event) {
			dialog.SetAttribute("open", "")
		}))

	}

}

func (view Tasks) Enter() bool {

	view.Refresh()

	return true

}

func (view Tasks) Leave() bool {
	return true
}

func (view Tasks) Refresh() {

	schema, err := actions.GetTasks(view.Main.Client)

	if err == nil {
		view.Schema.Tasks = schema.Tasks
		view.Main.Storage.Write("tasks", schema)
	}

	view.Render()

}

func (view Tasks) Render() {

	table := view.GetElement("table")

	if table != nil {

		html := ""
		ids  := make([]int, 0)

		for _, task := range view.Schema.Tasks {
			ids = append(ids, task.ID)
		}

		sort.Ints(ids)

		for i := 0; i < len(ids); i++ {

			task := view.Schema.Tasks[ids[i]]
			html += view.RenderTask(task)

		}

		tbody := table.QuerySelector("tbody")

		if tbody != nil {
			tbody.SetInnerHTML(html)
		}

	}

}

func (view Tasks) RenderTask(task *schemas.Task) string {

	var result string

	id := strconv.Itoa(task.ID)

	result += "<tr data-id=\"" + id + "\">"
	result += "<td>" + strconv.Itoa(task.ID) + "</td>"
	result += "<td>" + task.Title + "</td>"

	if task.Done == true {
		result += "<td><input type=\"checkbox\" checked /></td>"
	} else {
		result += "<td><input type=\"checkbox\" /></td>"
	}

	result += "</tr>"

	return result

}

func (view Tasks) CloseDialog() {

	dialog := view.GetElement("dialog")

	if dialog != nil {

		texts := dialog.QuerySelectorAll("input[type=\"text\"]")
		bools := dialog.QuerySelectorAll("input[type=\"checkbox\"]")

		for _, element := range texts {
			element.Value.Set("value", "")
		}

		for _, element := range bools {
			element.Value.Set("checked", false)
		}

		dialog.RemoveAttribute("open")

	}

}
