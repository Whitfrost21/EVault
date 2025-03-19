package login

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Whitfrost21/EVault/models"
)

var loginButton *widget.Button

// Function to create the login page
func CreateLoginPage(myWindow fyne.Window, navbar fyne.Widget, maincontent fyne.Widget) fyne.CanvasObject {
	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	loginButton = widget.NewButton("Login", func() {
		// username := usernameEntry.Text
		// password := passwordEntry.Text

		// if username == "admin" && password == "password123" {
		models.LogStatus = true
		myWindow.SetContent(container.NewBorder(navbar, nil, nil, nil, maincontent))
		myWindow.Resize(fyne.NewSize(1200, 600))
		// } else {
		// 	errorLabel := widget.NewLabel("Invalid credentials, please try again.")
		// 	content := container.NewVBox(widget.NewLabel("Username:"), usernameEntry, widget.NewLabel("Password:"), passwordEntry, loginButton, errorLabel)
		// 	myWindow.SetContent(content)
		// 	models.LogStatus = false
		// }
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
