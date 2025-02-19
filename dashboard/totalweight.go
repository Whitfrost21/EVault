package dashboard

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetWeight(url string) (float64, error) {

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var weight float64
	err = json.Unmarshal(body, &weight)
	if err != nil {
		return 0, err
	}

	return weight, nil
}

func DisplayWeight(content *fyne.Container, weight float64) {

	weightline := fmt.Sprintf("%.2f", weight)

	label := widget.NewLabel("Total Collected Waste:" + weightline + "kg")

	centeredcon := fyne.NewContainerWithLayout(
		layout.NewCenterLayout(),
		widget.NewLabel(label.Text),
	)

	content.Objects = []fyne.CanvasObject{centeredcon}
	content.Refresh()
}
