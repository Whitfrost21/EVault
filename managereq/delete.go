package managereq

import (
	"fmt"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Deleterequest(window fyne.Window) fyne.CanvasObject {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("Enter Request ID to Delete")

	return container.NewVBox(
		widget.NewLabel("Delete Request"),
		idEntry,
		widget.NewButton("Delete", func() {

			if idEntry.Text == "" {
				dialog.ShowInformation("Error", "Request ID is required!", window)
				return
			}

			apiURL := fmt.Sprintf("http://localhost:8080/pickuprequest/%s", idEntry.Text)

			req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to create DELETE request: %v", err), window)
				return
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to send DELETE request: %v", err), window)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "Success",
					Content: fmt.Sprintf("Request ID %s has been successfully deleted!", idEntry.Text),
				})

				idEntry.SetText("")
			} else {
				dialog.ShowError(fmt.Errorf("failed to delete request: received status code %d", resp.StatusCode), window)
			}
		}),
	)
}
