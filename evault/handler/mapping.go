package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/Whitfrost21/EVault/evault/models"
	"gorm.io/gorm"
)

var client = &http.Client{Timeout: 10 * time.Second}

func Getlocations(ctx context.Context) {

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("stopped collecting locations")
			return
		case <-ticker.C:
			var request models.Pickuprequest
			err := models.Db.Where("status = ?", false).Order("priority asc").First(&request).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println("No pending request... queue is empty")
				time.Sleep(5 * time.Minute)
				continue
			} else if err != nil {
				log.Printf("Error fetching next request: %v\n", err)
				time.Sleep(1 * time.Minute)
				continue
			}
			centerLat, centerLng := 16.825574, 74.608154 //change coords to center's coords
			lat, lon := request.Latitude, request.Longitude
			distance, Time, err := Getrouteinfo(lat, lon, centerLat, centerLng)
			if err != nil {
				log.Println("error while finding route", err)
				return
			}
			request.Distance = distance
			request.TravelTime = Time
			log.Printf("distance: %.2fkm and time:%.2fmins is calculated for the request: %d\n", distance, Time, request.Id)

			if err := models.Db.Save(&request).Error; err != nil {
				log.Printf("error while saving after calculating distance and time %d , %v", request.Id, err)
				time.Sleep(1 * time.Minute)
				continue
			}
			time.Sleep(1 * time.Minute)
		}
	}
}

func Getrouteinfo(lat1, lon1, lat2, lon2 float64) (float64, float64, error) {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("error while loading godotenv")
	// }
	// apikey := os.Getenv("API_KEY")
	apikey := " "
	url := fmt.Sprintf("https://graphhopper.com/api/1/route?point=%f,%f&point=%f,%f&vehicle=car&locale=en&key=%s", lat1, lon1, lat2, lon2, apikey)
	resp, err := client.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var ghresponse models.Graphopresponse
	if err := json.NewDecoder(resp.Body).Decode(&ghresponse); err != nil {
		return 0, 0, err
	}
	if len(ghresponse.Paths) == 0 {
		return 0, 0, fmt.Errorf("no route found")
	}
	distancekm := ghresponse.Paths[0].Distance / 1000
	timemin := float64(ghresponse.Paths[0].Time) / 60000
	return distancekm, timemin, nil
}
