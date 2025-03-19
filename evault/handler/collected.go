package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Whitfrost21/EVault/evault/models"
	"github.com/Whitfrost21/EVault/notifymesg"
	"github.com/gen2brain/beeep"
	"gorm.io/gorm"
)

func Decidequality(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("no pending requests , quality distribution stopped...")
			return
		default:
			var nextrequest models.Collectedrequests
			err := models.Db.Where("quality = ? OR quality IS NULL", "").Order("priority asc").First(&nextrequest).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println("No pending request, retrying in 5 minutes...")
				time.Sleep(5 * time.Minute)
				continue
			} else if err != nil {
				log.Printf("Error fetching request: %v\n", err)
				time.Sleep(1 * time.Minute)
				continue
			}

			//weight according to collected product
			weight := 12.32
			if err := models.Db.Model(&nextrequest).Update("weight", weight).Error; err != nil {
				log.Printf("error deciding weight:%d:%v\n", nextrequest.Id, err)
				time.Sleep(1 * time.Minute)
				continue
			}

			//  quality can be Good,Medium,Bad which is necessary to detemine cost
			quality := "Good"
			if err := models.Db.Model(&nextrequest).Update("quality", quality).Error; err != nil {
				log.Printf("error deciding quality:%d:%v\n", nextrequest.Id, err)
				time.Sleep(1 * time.Minute)
				continue
			}
			cost, err := Calculatecost(nextrequest, quality)
			if err != nil {
				log.Printf("error calculating the cost : %d,%v", nextrequest.Id, err)
				continue
			}
			if err := models.Db.Model(&nextrequest).Update("cost", cost).Error; err != nil {
				log.Printf("error while updating the cost for id:%d,%v", nextrequest.Id, err)
				continue
			}
			log.Printf("quality and cost decided for %d quality:%s cost:%.2f\n", nextrequest.Id, quality, cost)
			notifymesg.AddNotification(nextrequest.Name, fmt.Sprintf("we just evaulated quality for your waste which is :%s", quality))
			if err := SendNotification(nextrequest); err != nil {
				log.Printf("error sending notification:%d,%v", nextrequest.Id, err)
				continue
			}
			if err := Movetohistory(nextrequest); err != nil {
				log.Printf("error while storing in history for request:%d", nextrequest.Id)
			}
			time.Sleep(1 * time.Minute)
		}

	}
}

func Calculatecost(request models.Collectedrequests, quality string) (float64, error) { // calculate the cost by specified quality of product

	for _, rule := range models.CostRules {
		if rule.Quality == request.Quality && rule.Wastetype == request.Wastetype {
			return rule.Cost * float64(request.Quantity), nil
		}
	}
	return 0, fmt.Errorf("error quality or wastetype is not invalid for id:%d quality:%s wastetype:%s", request.Id, quality, request.Wastetype)
}

func SendNotification(request models.Collectedrequests) error {
	message := fmt.Sprintf(`Hello! Your E-waste has been successfully collected and recycled in Evault.
Following are your overall details:
- Name: %s
- Collected Waste: %s
- Quantity: %d
- Weight: %.2f
- Cost: %.2f

Thank you!`, request.Name, request.Wastetype, request.Quantity, request.Weight, request.Cost)

	log.Printf("Sending in-app notification to: %s", request.Name)

	notifymesg.AddNotification(request.Name, message)
	err := beeep.Notify("E-Waste Collection Completed", message, "")
	if err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}

	log.Println("In-app notification sent successfully for ID:", request.Id)
	return nil
}

// after notification move completed requests to history
func Movetohistory(nextrequest models.Collectedrequests) error {
	var completedreq models.Collectedrequests
	log.Printf("Completed request: %v \n", nextrequest.Id)
	if err := models.Db.Where("id=?", nextrequest.Id).First(&completedreq).Error; err != nil {
		log.Printf("error fetching request %v\n", nextrequest.Id)
		return err
	}
	collected := models.History{
		Name:        nextrequest.Name,
		Address:     nextrequest.Address,
		Latitude:    nextrequest.Latitude,
		Longitude:   nextrequest.Longitude,
		Wastetype:   nextrequest.Wastetype,
		Description: nextrequest.Description,
		Phone:       nextrequest.Phone,
		Quantity:    nextrequest.Quantity,
		Cost:        nextrequest.Cost,
		Quality:     nextrequest.Quality,
		Weight:      nextrequest.Weight,
		CompletedAt: nextrequest.CompletedAt,
	}
	if err := models.Db.Create(&collected).Error; err != nil {
		log.Printf("error while inserting collection: %d %v\n", nextrequest.Id, err)
		return err
	}
	if err := models.Db.Delete(&nextrequest).Error; err != nil {
		log.Printf("error while deleting record %v \n", nextrequest.Id)
		return err
	}
	log.Printf("successfully stored request:%d to history.", nextrequest.Id)
	return nil
}
