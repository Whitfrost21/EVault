package managereq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Whitfrost21/EVault/dashboard/mapview"
	"github.com/Whitfrost21/EVault/models"
)

func CreateUpdateForm(window fyne.Window) fyne.CanvasObject {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("Enter Request ID")

	return container.NewVBox(
		widget.NewLabel("Update Request"),
		idEntry,
		widget.NewButton("Fetch and Update", func() {
			if idEntry.Text == "" {
				dialog.ShowInformation("Error", "Request ID is required!", window)
				return
			}
			nameEntry := widget.NewEntry()
			phoneEntry := widget.NewEntry()
			wasteTypeSelect := widget.NewSelect(models.Wastetypes, nil)
			description := widget.NewEntry()
			description.SetPlaceHolder("(Optional:if wastetype is other)")
			quantityEntry := widget.NewEntry()
			openMapButton := widget.NewButton("Select Location on Map", func() {

				cmd := exec.Command("firefox", "http://localhost:8081/map")
				err := cmd.Start()
				if err != nil {
					log.Printf("error while opening the firefox: %v", err)
				}
			})
			form := widget.NewForm(
				widget.NewFormItem("Name", nameEntry),
				widget.NewFormItem("Phone No.", phoneEntry),
				widget.NewFormItem("Waste Type", wasteTypeSelect),
				widget.NewFormItem("Description", description),
				widget.NewFormItem("", openMapButton),
				widget.NewFormItem("Quantity", quantityEntry),
			)

			// Handle form submission
			form.OnSubmit = func() {
				// Ensure all fields are filled
				if idEntry.Text == "" || nameEntry.Text == "" || phoneEntry.Text == "" || wasteTypeSelect.Selected == "" || quantityEntry.Text == "" {
					dialog.ShowInformation("Error", "All fields are required!", window)
					return
				}
				quantity, err := strconv.Atoi(quantityEntry.Text)
				if err != nil {
					log.Printf("error while converting quantity to int %v", err)
					return
				}

				data := models.Sendreq{
					Name:        nameEntry.Text,
					Latitude:    mapview.Lat,
					Longitude:   mapview.Lng,
					Phone:       phoneEntry.Text,
					Description: description.Text,
					Wastetype:   wasteTypeSelect.Selected,
					Quantity:    quantity,
				}
				jsonData, err := json.Marshal(data)
				if err != nil {
					dialog.ShowError(fmt.Errorf("failed to create JSON payload: %v", err), window)
					return
				}

				apiURL := fmt.Sprintf("http://localhost:8080/pickuprequest/%s", idEntry.Text)

				req, err := http.NewRequest(http.MethodPut, apiURL, bytes.NewBuffer(jsonData))
				if err != nil {
					dialog.ShowError(fmt.Errorf("failed to create PUT request: %v", err), window)
					return
				}
				req.Header.Set("Content-Type", "application/json")

				// Send the PUT request using http.Client
				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					dialog.ShowError(fmt.Errorf("failed to send request: %v", err), window)
					return
				}
				defer resp.Body.Close()

				// Handle response
				if resp.StatusCode == http.StatusOK {
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title:   "Success",
						Content: "Your request has been updated successfully!",
					})
				} else {
					dialog.ShowError(fmt.Errorf("failed to update request: received status code %d", resp.StatusCode), window)
				}
			}

			// Show the form in a dialog or a container
			dialog.ShowCustom("Update Request", "Close", container.NewVBox(form), window)
		}),
	)
}
