package dashboard

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type BarChart struct {
	data         []float32
	labels       []string
	title        string
	heightFactor float32
}

func CreateChart(b *BarChart) fyne.CanvasObject {

	var bars []fyne.CanvasObject

	barWidth := float32(50)   // Reduced Bar width
	barSpacing := float32(40) // Increased spacing between bars
	maxHeight := float32(150) // Reduced Max height of the chart container

	if b.heightFactor == 0 {
		b.heightFactor = 1.0 // Default scaling factor is 1
	}

	var maxDataValue float32
	for _, value := range b.data {
		if value > maxDataValue {
			maxDataValue = value
		}
	}

	barColors := []color.Color{
		// color.RGBA{R: 255, G: 87, B: 34, A: 255},  // Red
		color.RGBA{R: 33, G: 150, B: 243, A: 255}, // Blue
		// color.RGBA{R: 76, G: 175, B: 80, A: 255},  // Green
		// color.RGBA{R: 255, G: 235, B: 59, A: 255}, // Yellow
		// color.RGBA{R: 156, G: 39, B: 176, A: 255}, // Purple
	}

	// Create bars for the chart with custom height scaling
	for i, value := range b.data {

		barHeight := 4 * value
		bar := canvas.NewRectangle(barColors[i%len(barColors)])
		bar.Resize(fyne.NewSize(barWidth, barHeight))
		bar.Move(fyne.NewPos(float32(i)*(barWidth+barSpacing), maxHeight-barHeight))

		// Add labels with values above the bars (with a larger gap between bar and label)
		label := widget.NewLabel(fmt.Sprintf("%.1f", value))
		label.Move(fyne.NewPos(float32(i)*(barWidth+barSpacing), maxHeight-barHeight-60)) // Increased gap between label and bar

		itemLabel := b.labels[i]
		var itemLabels []*widget.Label
		if itemLabel == "Charging Cables" {
			itemLabels = append(itemLabels, widget.NewLabel("Charging"))
			itemLabels = append(itemLabels, widget.NewLabel("Cables"))
		} else if itemLabel == "Scrap Metal" {
			itemLabels = append(itemLabels, widget.NewLabel("Scrap"))
			itemLabels = append(itemLabels, widget.NewLabel("Metal"))
		} else if itemLabel == "Lithium Cells" {
			itemLabels = append(itemLabels, widget.NewLabel("Lithium"))
			itemLabels = append(itemLabels, widget.NewLabel("Cells"))
		} else if itemLabel == "Plastic Casings" {
			itemLabels = append(itemLabels, widget.NewLabel("Plastic"))
			itemLabels = append(itemLabels, widget.NewLabel("Casings"))
		} else if itemLabel == "Circuit Boards" {
			itemLabels = append(itemLabels, widget.NewLabel("Circuit"))
			itemLabels = append(itemLabels, widget.NewLabel("Boards"))
		} else if itemLabel == "Battery Chips" {
			itemLabels = append(itemLabels, widget.NewLabel("Battery"))
			itemLabels = append(itemLabels, widget.NewLabel("Chips"))
		} else if itemLabel == "Worn Brushes" {
			itemLabels = append(itemLabels, widget.NewLabel("Worn"))
			itemLabels = append(itemLabels, widget.NewLabel("Brushes"))
		} else if itemLabel == "Charging ports" {
			itemLabels = append(itemLabels, widget.NewLabel("Charging"))
			itemLabels = append(itemLabels, widget.NewLabel("ports"))
		} else if itemLabel == "Outdated Chips" {
			itemLabels = append(itemLabels, widget.NewLabel("Outdated"))
			itemLabels = append(itemLabels, widget.NewLabel("Chips"))
		} else if itemLabel == "Plastic Trim" {
			itemLabels = append(itemLabels, widget.NewLabel("Plastic"))
			itemLabels = append(itemLabels, widget.NewLabel("Trim"))
		} else {
			itemLabels = append(itemLabels, widget.NewLabel(itemLabel))
		}

		for j, item := range itemLabels {
			item.Move(fyne.NewPos(float32(i)*(barWidth+barSpacing), maxHeight+60+float32(j)*25)) // Adjust positioning for multi-line
			bars = append(bars, item)
		}

		bars = append(bars, bar, label)
	}

	// Create title label
	titleLabel := widget.NewLabel(b.title)
	titleLabel.Move(fyne.NewPos(0, -100)) // Position title well above the bars to avoid overlap

	// Create a container for the bars and labels
	barContainer := container.NewWithoutLayout(bars...)
	// totalWidth := float32(len(b.data)) * (barWidth + barSpacing)
	// scroll := container.NewHScroll(con)
	// Return the container with title and bars (wrapped in the scroll container for  scrolling)
	return container.NewVBox(
		widget.NewLabel("Waste Collection Graph:"),
		// Title label
		barContainer,
	)
}
