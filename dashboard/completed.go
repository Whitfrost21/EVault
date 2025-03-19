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
type Collectedrequests struct {
	Id          int     `gorm:"primaryKey"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	WasteType   string  `json:"wastetype"`
	Description string  `json:"description"`
	Phone       string  `json:"phone"`
	Quantity    int     `json:"quantity"`
	Priority    int     `json:"priority"`
	Status      bool    `json:"status"`
	Quality     string  `json:"quality"`
	Cost        int     `json:"cost"`
	CompletedAt string  `json:"completedat"`
}

// FetchHistory fetches history data via an HTTP GET request
func Fetchcompleted(url string) ([]Collectedrequests, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var requests []Collectedrequests
	err = json.Unmarshal(body, &requests)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

// create a table and also add a scroll bar to it
func Displaycompleted(dynamicContent *fyne.Container, requests []Collectedrequests) {

	headers := []string{"ID", "Name", "Address", "Waste-Type", "Description", "Latitude", "Longitude", "Phone", "Quantity", "Priotiry", "Status", "Quality", "Cost", "CompletedAt"}

	// Data for the table
	data := make([][]string, len(requests)+1)
	data[0] = headers
	for i, req := range requests {
		data[i+1] = []string{
			fmt.Sprintf("%d", req.Id),
			req.Name,
			req.Address,
			req.WasteType,
			req.Description,
			fmt.Sprintf("%f", req.Latitude),
			fmt.Sprintf("%f", req.Longitude),
			req.Phone,
			fmt.Sprintf("%d", req.Quantity),
			fmt.Sprintf("%d", req.Priority),
			fmt.Sprintf("%v", req.Status),
			req.Quality,
			fmt.Sprintf("%d", req.Cost),
			req.CompletedAt,
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
	table.SetColumnWidth(2, 300)  //Address
	table.SetColumnWidth(3, 200)  // Waste Type
	table.SetColumnWidth(4, 300)  // Description
	table.SetColumnWidth(5, 100)  // latitude
	table.SetColumnWidth(6, 100)  // Longitude
	table.SetColumnWidth(7, 100)  // phone
	table.SetColumnWidth(8, 80)   // Quantity
	table.SetColumnWidth(9, 50)   //priority
	table.SetColumnWidth(10, 100) //status
	table.SetColumnWidth(11, 100) //quality
	table.SetColumnWidth(12, 80)  //cost
	table.SetColumnWidth(13, 300) //completedAt

	// Wrap the table in a scroll container
	scroll := container.NewVScroll(table)
	scroll.SetMinSize(fyne.NewSize(800, 400)) // Set desired size for the scroll area

	// Replace the content with the scroll container
	dynamicContent.Objects = []fyne.CanvasObject{scroll}
	dynamicContent.Refresh()
}
