package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateThemeswap(app fyne.App) fyne.CanvasObject {

	islightmode := false
	var togglebutton *widget.Button
	togglebutton = widget.NewButton("light mode", func() {
		if islightmode {
			app.Settings().SetTheme(theme.DarkTheme())
			togglebutton.SetText("light mode")
		} else {
			app.Settings().SetTheme(theme.LightTheme())
			togglebutton.SetText("dark mode")
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
