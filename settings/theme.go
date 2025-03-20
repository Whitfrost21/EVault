package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var islightmode = false

func CreateThemeswap(app fyne.App) fyne.CanvasObject {
	var togglebutton *widget.Button

	togglebutton = widget.NewButton("Change Theme", func() {
		if islightmode {
			app.Settings().SetTheme(theme.DarkTheme())

			islightmode = true
		} else {
			app.Settings().SetTheme(theme.LightTheme())

		}
		islightmode = !islightmode
	})

	content := container.NewVBox(
		widget.NewLabel("Settings"),
		widget.NewLabel("Theme:"),
		togglebutton,
	)

	return content
}
