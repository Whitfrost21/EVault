package login

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Whitfrost21/EVault/models"
	"github.com/gen2brain/beeep"
)

var loginButton *widget.Button

// Function to create the login page
func CreateLoginPage(myWindow fyne.Window, navbar fyne.Widget, maincontent fyne.Widget) fyne.CanvasObject {
	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	loginButton = widget.NewButton("Login", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text
		if username == "" || password == "" {
			beeep.Notify("EVault", "Enter valid username and password pls", "")
		} else if username != "admin" && password != "password123" {
			beeep.Notify("EVault", "Username or password is incorrect", "")
			usernameEntry.SetText("")
			passwordEntry.SetText("")
		} else {
			models.LogStatus = true
			myWindow.SetContent(container.NewBorder(navbar, nil, nil, nil, maincontent))
			myWindow.Resize(fyne.NewSize(1200, 600))
		}
	})
	content := container.NewVBox(
		widget.NewLabelWithStyle("Welcome to EVault", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabel("Username:"),
		usernameEntry,
		widget.NewLabel("Password:"),
		passwordEntry,
		loginButton,
	)

	return content
}
