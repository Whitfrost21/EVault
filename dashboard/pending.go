package dashboard

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// WasteRequest represents a single request's details
type Pickuprequest struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	WasteType   string  `json:"wastetype"`
	Description string  `json:"description"`
	Phone       string  ` json:"phone"`
	Quantity    int     ` json:"quantity"`
	Status      bool    `json:"status"`
	Priority    int     `json:"priority"`
	Distance    float64 `json:"distance"`
	TravelTime  float64 `json:"traveltime"`
}

// FetchHistory fetches history data via an HTTP GET request
func Fetchpending(url string) ([]Pickuprequest, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var requests []Pickuprequest
	err = json.Unmarshal(body, &requests)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

// create a table and also add a scroll bar to it
func Displaypending(dynamicContent *fyne.Container, requests []Pickuprequest) {

	headers := []string{"ID", "Name", "Latitude", "Longitude", "WasteType", "Description", "Phone", "Quantity", "Status", "priority", "Distance", "TravelTime"}

	// Data for the table
	data := make([][]string, len(requests)+1)
	data[0] = headers
	for i, req := range requests {
		data[i+1] = []string{
			fmt.Sprintf("%d", req.Id),
			req.Name,
			fmt.Sprintf("%f", req.Latitude),
			fmt.Sprintf("%f", req.Longitude),
			req.WasteType,
			req.Description,
			req.Phone,
			fmt.Sprintf("%d", req.Quantity),
			fmt.Sprintf("%v", req.Status),
			fmt.Sprintf("%d", req.Priority),
			fmt.Sprintf("%f", req.Distance),
			fmt.Sprintf("%f", req.TravelTime),
		}
	}

	// Create the table widget
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(headers)
		},
		func() fyne.CanvasObject {
			// Create cells with a fixed size
			label := widget.NewLabel("")
			label.Wrapping = fyne.TextWrapWord
			return label
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			// Populate table cells with data
			cell.(*widget.Label).SetText(data[id.Row][id.Col])
		},
	)

	// Set minimum size for the table to avoid collapsing
	table.SetColumnWidth(0, 50)   // ID
	table.SetColumnWidth(1, 200)  // Name
	table.SetColumnWidth(2, 100)  // latitude
	table.SetColumnWidth(3, 100)  // longitude
	table.SetColumnWidth(4, 200)  // wastetype
	table.SetColumnWidth(5, 300)  // description
	table.SetColumnWidth(6, 100)  // phone
	table.SetColumnWidth(7, 80)   // Quantity
	table.SetColumnWidth(8, 80)   //status
	table.SetColumnWidth(9, 100)  //priority
	table.SetColumnWidth(10, 100) //distance
	table.SetColumnWidth(11, 120) //travel time

	// Wrap the table in a scroll container
	scroll := container.NewVScroll(table)
	scroll.SetMinSize(fyne.NewSize(800, 400)) // Set desired size for the scroll area

	// Replace the content with the scroll container
	dynamicContent.Objects = []fyne.CanvasObject{scroll}
	dynamicContent.Refresh()
}
