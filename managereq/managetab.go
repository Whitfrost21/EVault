package managereq

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Createmanagetab(window fyne.Window) fyne.CanvasObject {

	titlelabel := widget.NewLabelWithStyle("Manage Requests", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	dynamiccontent := container.NewVBox()

	updatecontent := func(newcontent fyne.CanvasObject) {
		dynamiccontent.Objects = []fyne.CanvasObject{newcontent}
		dynamiccontent.Refresh()
	}

	buttons := container.NewVBox(
		widget.NewButton("Update Request", func() {
			updatecontent(CreateUpdateForm(window))
		}),
		widget.NewButton("Delete Request", func() {
			updatecontent(Deleterequest(window))
		}),
		widget.NewButton("Get Request", func() {
			updatecontent(Getrequest(window))
		}),
	)
	mainContent := container.NewBorder(
		container.NewVBox(
			titlelabel,
			widget.NewSeparator(),
			buttons,
		),
		nil, nil, nil, dynamiccontent,
	)
	return mainContent
}
