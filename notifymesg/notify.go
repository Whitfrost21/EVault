package notifymesg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Notification struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Read    bool   `json:"read"`
}

var notifications []Notification
var unreadCount int
var filePath = "notifications.json"
var notificationButtons []*widget.Button

func loadNotifications() error {
	file, err := os.Open(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	if err == nil {
		byteValue, _ := ioutil.ReadAll(file)
		err = json.Unmarshal(byteValue, &notifications)
		if err != nil {
			return fmt.Errorf("error unmarshalling JSON: %v", err)
		}
	}
	return nil
}

func saveNotifications() error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	byteValue, err := json.MarshalIndent(notifications, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	_, err = file.Write(byteValue)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

func updateUnreadCount() int {
	unreadCount = 0
	for _, n := range notifications {
		if !n.Read {
			unreadCount++
		}
	}
	return unreadCount
}

func markAsRead(index int) {
	if index >= 0 && index < len(notifications) {
		notifications[index].Read = true
		updateUnreadCount()
		saveNotifications()
		updateButtonColors()
	}
}

func AddNotification(name, message string) {
	notification := Notification{
		Name:    name,
		Message: message,
		Read:    false,
	}
	notifications = append(notifications, notification)
	saveNotifications()
	updateButtonColors()
}

func Createnotificationtab(window fyne.Window) fyne.CanvasObject {
	notificationContainer := container.NewVBox()

	err := loadNotifications()
	if err != nil {
		fmt.Println("Error loading notifications:", err)
	}

	notificationContainer.Objects = nil
	notificationButtons = nil

	sort.Slice(notifications, func(i, j int) bool {
		return !notifications[i].Read && notifications[j].Read
	})
	for i, n := range notifications {

		status := "Unread"
		if n.Read {
			status = "Read"
		}
		button := widget.NewButton(fmt.Sprintf("%s: (%s)", n.Name, status), func() {
			showFullNotification(window, n.Message, i)
		})

		updateButtonAppearance(button, n.Read)

		notificationButtons = append(notificationButtons, button)
		notificationContainer.Add(button)
	}

	return notificationContainer
}

func updateButtonAppearance(button *widget.Button, read bool) {
	if read {
		button.Importance = widget.LowImportance
	} else {
		button.Importance = widget.HighImportance
	}
}

func updateButtonColors() {
	for i, button := range notificationButtons {
		updateButtonAppearance(button, notifications[i].Read)
	}
}

func showFullNotification(window fyne.Window, message string, index int) {
	infoWindow := fyne.CurrentApp().NewWindow("Notification Details")

	content := container.NewVBox(
		widget.NewLabel(message),
		widget.NewButton("Mark as Read", func() {
			markAsRead(index)
			infoWindow.Close()
			window.Content().Refresh()
		}),
	)
	infoWindow.SetContent(content)
	infoWindow.Resize(fyne.NewSize(600, 250))
	infoWindow.Show()
}
