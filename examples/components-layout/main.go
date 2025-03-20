package main

import gooey "github.com/cookiengineer/gooey/pkg"
import "github.com/cookiengineer/gooey/pkg/console"
import "github.com/cookiengineer/gooey-components/pkg/components"
import "github.com/cookiengineer/gooey-components/pkg/interfaces"
import "time"

func main() {

	// header := components.NewHeader(gooey.Document.QuerySelector("header"))
	footer := components.ToFooter(gooey.Document.QuerySelector("footer"))

	button_cancel  := components.NewButton("Cancel", "cancel")
	button_confirm := components.NewButton("Confirm", "confirm")

	footer.SetLayout([]interfaces.Component{
		&button_cancel,
	}, []interfaces.Component{
	}, []interfaces.Component{
		&button_confirm,
	})

	console.Log(footer)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
