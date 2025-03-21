package dashboard

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateDashboard(window fyne.Window) fyne.CanvasObject {

	titleLabel := canvas.NewText("Evault Dashboard", color.White)
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Motivational message
	motivationLabel := canvas.NewText("Recycle your e-waste and save the planet!", color.White)
	motivationLabel.Alignment = fyne.TextAlignCenter

	data1 := []float32{1, 6, 2, 1, 4, 3, 8, 5, 3, 5, 6, 7}
	// data2 := []float32{3, 2, 3, 5, 6, 8, 2, 4, 4, 2, 6}
	label1 := []string{
		"Lithium Cells", "Batteries", "Connectors", "Charging Cables",
		"Scrap Metal", "Air Filters", "LED lights", "Wires", "Chips", "Outdated Chips", "Frames", "Plastic Trim", "Compressor"} //
	// label2 := []string{"Plastic Casings", "Circuit Boards", " Battery Chips", "Electrolyte", "Rotor", "Stator", "Worn Brushes", "Charging ports", "Tires", "Rims", "Dashboards"}
	//
	// Create the bar chart with custom height scaling factor (adjust as needed)
	chart1 := &BarChart{data: data1, labels: label1, title: "E-Waste Collection", heightFactor: 1.5}

	// chart2 := &BarChart{data: data2, labels: label2, title: "E-Waste Collection", heightFactor: 1.5}
	bar1 := CreateChart(chart1)
	// bar2 := CreateChart(chart2)
	// Add the chart to the container

	chartContainer := container.New(
		layout.NewVBoxLayout(),
		container.New(layout.NewHBoxLayout(), canvas.NewText("Collect Your E-waste Now", color.White)),
	)
	bar1.Resize(fyne.NewSize(500, 500))
	// bar2.Resize(fyne.NewSize(500, 500))
	// barbox := container.NewMax(bar1, bar2)
	// scrollbox := container.NewVScroll(barbox)
	// scrollbox.Resize(fyne.NewSize(1000, 800))
	separator := widget.NewSeparator()
	dynamicContent := container.NewVBox(
		titleLabel,
		separator,
		separator,
		separator,
		separator,
		chartContainer,
		separator,
		separator,
		separator,
		separator,
		separator,
		separator,
		separator,
		separator,
		separator, separator,
		separator,
		separator,
		separator,
		separator,
		separator,
		separator, separator,
	)

	backgroundImage := canvas.NewImageFromFile("/home/pz/Downloads/Ev.jpg")
	backgroundImage.FillMode = canvas.ImageFillStretch
	backgroundImage.Resize(fyne.NewSize(600, 500))

	contentWithBackground := container.NewMax(backgroundImage, dynamicContent)

	// Buttons section - Grouped in a card-like layout
	buttons := container.NewVBox(
		widget.NewButtonWithIcon("Add New Request", theme.ContentAddIcon(), func() {
			dynamicContent.Objects = []fyne.CanvasObject{
				Showrequestform(window),
			}
			backgroundImage.Hide()
			dynamicContent.Refresh()
			bar1.Hide()
			// bar2.Hide()
		}),
		widget.NewButtonWithIcon("Collected Requests", theme.FileApplicationIcon(), func() {
			url := "http://localhost:8080/getcollectedlist"
			requests, err := Fetchcompleted(url)
			if err != nil {
				dynamicContent.Objects = []fyne.CanvasObject{
					widget.NewLabel(fmt.Sprintf("Error: %v", err)),
				}
				backgroundImage.Hide()
				dynamicContent.Refresh()
				return
			}
			backgroundImage.Hide()
			Displaycompleted(dynamicContent, requests)
		}),
		widget.NewButtonWithIcon("Pending Requests", theme.HelpIcon(), func() {
			url := "http://localhost:8080/pickuprequest"
			requests, err := Fetchpending(url)
			if err != nil {
				dynamicContent.Objects = []fyne.CanvasObject{
					widget.NewLabel(fmt.Sprintf("Error: %v", err)),
				}
				backgroundImage.Hide()
				dynamicContent.Refresh()
				return
			}
			backgroundImage.Hide()
			Displaypending(dynamicContent, requests)
		}),
		widget.NewButtonWithIcon("History", theme.SearchIcon(), func() {
			url := "http://localhost:8080/gethistory"
			requests, err := FetchHistory(url)
			if err != nil {
				dynamicContent.Objects = []fyne.CanvasObject{
					widget.NewLabel(fmt.Sprintf("Error: %v", err)),
				}
				backgroundImage.Hide()
				dynamicContent.Refresh()
				return
			}
			backgroundImage.Hide()
			DisplayHistory(dynamicContent, requests)
		}),
		widget.NewButton("Total Collection", func() {

			url := "http://localhost:8080/gettotalweight"
			weight, err := GetWeight(url)
			if err != nil {
				dynamicContent.Objects = []fyne.CanvasObject{
					widget.NewLabel(fmt.Sprintf("Error: %v", err)),
				}
				backgroundImage.Hide()
				dynamicContent.Refresh()
				return
			}
			backgroundImage.Hide()
			DisplayWeight(dynamicContent, weight)
			bar1.Hide()
			// bar2.Hide()
		}),
	)

	for _, button := range buttons.Objects {
		if b, ok := button.(*widget.Button); ok {
			b.Importance = widget.MediumImportance
			b.Resize(fyne.NewSize(200, 60)) // Resize for better UX
		}
	}

	dashboard := container.NewVBox(
		contentWithBackground,
		buttons,
		bar1,
	)
	dashboard.Resize(fyne.NewSize(800, 2000))
	return dashboard
}
