package dashboard

import (
	"Source/dashboard/mapview"
	"Source/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gen2brain/beeep"
	webview "github.com/webview/webview_go"
)

func Showrequestform(window fyne.Window) fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	phoneEntry := widget.NewEntry()
	wasteType := widget.NewSelect(models.Wastetypes, nil)
	quantityEntry := widget.NewEntry()
	// selectedCoords := widget.NewLabel("Selected Coordinates: None")

	formSubmitButton := widget.NewButton("Submit", func() {
		quantity, err := strconv.Atoi(quantityEntry.Text)
		if err != nil {
			log.Println("Error parsing quantity:", err)
		}

		data := models.Sendreq{
			Name:      nameEntry.Text,
			Latitude:  mapview.Lat,
			Longitude: mapview.Lng,
			Phone:     phoneEntry.Text,
			Wastetype: wasteType.Selected,
			Quantity:  quantity,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return
		}

		url := "http://localhost:8080/pickuprequest/"
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		fmt.Println("Response Status:", resp.Status)
		fmt.Println("request submitted!")
		err = beeep.Notify("Evault", "Request Submitted!!", "")
		if err != nil {
			log.Println("error sending beeep notification", err)
			return
		}

	})

	var w webview.WebView

	openMapButton := widget.NewButton("Select Location on Map", func() {
		if w != nil {
			log.Print("webveiw is not empty")
			w.Destroy()
		}

		w = webview.New(false)
		w.SetTitle("Select a Location")
		w.SetSize(800, 600, 0)
		defer w.Destroy()
		w.Navigate("http://localhost:8081/map")

		w.Run()
		// lat := strconv.FormatFloat(mapview.Lat, 'f', -1, 64)
		// lng := strconv.FormatFloat(mapview.Lng, 'f', -1, 64)
		// selectedCoords.SetText(fmt.Sprintf("Selected Coordinates: %s , %s", lat, lng))
	})

	form := widget.NewForm(
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Phone No.", phoneEntry),
		widget.NewFormItem("Waste Type", wasteType),
		widget.NewFormItem("Quantity", quantityEntry),
		widget.NewFormItem("", openMapButton),
	)

	return container.NewVBox(
		widget.NewLabelWithStyle("Add New Request", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		form,
		formSubmitButton,
		// selectedCoords,
	)
}
