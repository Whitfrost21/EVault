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
type WasteRequest struct {
	Id          int     `json:"Id"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	WasteType   string  `json:"wastetype"`
	Description string  `json:"description"`
	Phone       string  `json:"phone"`
	Quantity    int     `json:"quantity"`
	Cost        int     `json:"Cost"`
	Quality     string  `json:"quality"`
	CompletedAt string  `json:"completedat"`
}

// FetchHistory fetches history data via an HTTP GET request
func FetchHistory(url string) ([]WasteRequest, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var requests []WasteRequest
	err = json.Unmarshal(body, &requests)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

// create a table and also add a scroll bar to it
func DisplayHistory(dynamicContent *fyne.Container, requests []WasteRequest) {

	headers := []string{"ID", "Name", "Waste-Type", "Phone", "Quantity", "Cost", "Quality", "Completed At"}

	// Data for the table
	data := make([][]string, len(requests)+1)
	data[0] = headers
	for i, req := range requests {
		data[i+1] = []string{
			fmt.Sprintf("%d", req.Id),
			req.Name,
			req.WasteType,
			req.Phone,
			fmt.Sprintf("%d", req.Quantity),
			fmt.Sprintf("%d", req.Cost),
			req.Quality,
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
	table.SetColumnWidth(0, 50)  // ID
	table.SetColumnWidth(1, 200) // Name
	table.SetColumnWidth(2, 200) // Waste Type
	table.SetColumnWidth(3, 100) // Phone
	table.SetColumnWidth(4, 80)  // Quantity
	table.SetColumnWidth(5, 80)  // Cost
	table.SetColumnWidth(6, 100) // Quality
	table.SetColumnWidth(7, 150) // Completed At

	// Wrap the table in a scroll container
	scroll := container.NewVScroll(table)
	scroll.SetMinSize(fyne.NewSize(800, 400)) // Set desired size for the scroll area

	// Replace the content with the scroll container
	dynamicContent.Objects = []fyne.CanvasObject{scroll}
	dynamicContent.Refresh()
}
