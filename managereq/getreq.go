package managereq

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Getrequest(window fyne.Window) fyne.CanvasObject {

	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("Enter Request ID")

	tableContainer := container.NewVBox(widget.NewLabel("Request Details will appear here."))

	idForm := container.NewVBox(
		widget.NewLabel("Get Request"),
		idEntry,
		widget.NewButton("Submit", func() {

			if idEntry.Text == "" {
				dialog.ShowInformation("Error", "Request ID is required!", window)
				return
			}

			apiURL := fmt.Sprintf("http://localhost:8080/pickuprequest/%s", idEntry.Text)

			resp, err := http.Get(apiURL)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to fetch request: %v", err), window)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				dialog.ShowError(fmt.Errorf("failed to fetch request: received status code %d", resp.StatusCode), window)
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to read response body: %v", err), window)
				return
			}

			var requestData map[string]interface{}
			err = json.Unmarshal(body, &requestData)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to parse JSON response: %v", err), window)
				return
			}

			tableContainer.Objects = []fyne.CanvasObject{}

			rows := make([]fyne.CanvasObject, 0)
			for key, value := range requestData {
				rows = append(rows, container.NewHBox(
					widget.NewLabelWithStyle(fmt.Sprintf("%s:", key), fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
					widget.NewLabel(fmt.Sprintf("%v", value)),
				))
			}

			tableContainer.Objects = rows
			tableContainer.Refresh()
		}),
	)

	mainContent := container.NewVBox(idForm, widget.NewSeparator(), tableContainer)

	return mainContent
}
