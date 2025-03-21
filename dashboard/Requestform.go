package dashboard

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
	"fyne.io/fyne/v2/widget"

	"github.com/Whitfrost21/EVault/dashboard/mapview"
	"github.com/Whitfrost21/EVault/models"
	"github.com/gen2brain/beeep"
)

var address string
var Lat, Lng float64

func Showrequestform(window fyne.Window) fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	phoneEntry := widget.NewEntry()
	wasteType := widget.NewSelect(models.Wastetypes, nil)
	description := widget.NewEntry()
	description.SetPlaceHolder("(Optional:only is wastetype is other)")
	quantityEntry := widget.NewEntry()

	formSubmitButton := widget.NewButton("Submit", func() {

		if nameEntry.Text == "" {
			err := beeep.Notify("Evault", "name cannot be empty", "")
			if err != nil {
				log.Printf("error %v", err)
			}
			return
		}
		if len(phoneEntry.Text) != 10 {
			err := beeep.Notify("Evault", "phone no must be 10 digits", "")
			if err != nil {
				log.Printf("error %v", err)
			}
			return
		}
		if wasteType.Selected == "" {
			err := beeep.Notify("Evault", "select the waste type pls", "")
			if err != nil {
				log.Printf("error %v", err)
			}
			return
		}
		if quantityEntry.Text == "" || quantityEntry.Text == "0" {
			err := beeep.Notify("Evault", "qunatity cannot be empty", "")
			if err != nil {
				log.Printf("error %v", err)
			}
			return
		}

		quantity, err := strconv.Atoi(quantityEntry.Text)
		if err != nil {
			log.Println("Error parsing quantity:", err)
		}
		Lat = mapview.Lat
		Lng = mapview.Lng
		// parseaddress here....
		address, err = GetAddress(Lat, Lng)
		if err != nil {
			log.Println("error while finding the address from lat/long:", err)
		}
		if address == "" {
			err := beeep.Notify("Evault", "Address cannot be empty", "")
			if err != nil {
				log.Printf("error %v", err)
			}
			return
		}

		data := models.Sendreq{
			Name:      nameEntry.Text,
			Address:   address,
			Latitude:  Lat,
			Longitude: Lng,
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

	// var w webview.WebView (just ignore this thing)

	openMapButton := widget.NewButton("Select Location on Map", func() {

		cmd := exec.Command("firefox", "http://localhost:8081/map")
		err := cmd.Start()
		if err != nil {
			log.Printf("error while opening the firefox: %v", err)
		}
		// lat := strconv.FormatFloat(mapview.Lat, 'f', -1, 64)
		// lng := strconv.FormatFloat(mapview.Lng, 'f', -1, 64)
		// selectedCoords.SetText(fmt.Sprintf("Selected Address: %s ", address))

		//this was my old configuration i used webveiw_go framework but it was causing some issues while running so i decided to use a straitforward way instead
		// Check if there is an existing instance, and safely terminate and destroy it
		// if w != nil {
		// 	r := webview.New(false)
		// 	r.Navigate("http://localhost:8081/map") // Set w to nil after destroying to avoid using an invalid instance
		// 	r.Run()
		// 	defer r.Destroy()
		// }

		// Create a new instance
		// if w == nil {
		// 	w = webview.New(false)
		// 	w.SetTitle("Select a Location")
		// 	w.SetSize(800, 600, 0)

		// Safely run the new instance
		// 	go func() {
		// 		w.Navigate("http://localhost:8081/map")
		// 		w.Run()
		// 		defer w.Destroy()
		// 	}()
		// }

	})

	form := widget.NewForm(
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Phone No.", phoneEntry),
		widget.NewFormItem("Waste Type", wasteType),
		widget.NewFormItem("Description", description),
		widget.NewFormItem("Quantity", quantityEntry),
		widget.NewFormItem("", openMapButton),
	)
	return container.NewVBox(
		widget.NewLabelWithStyle("Add New Request", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		form,
		formSubmitButton,
	)
}
