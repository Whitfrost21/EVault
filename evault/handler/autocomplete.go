package handler

import (
	"Source/evault/models"
	"Source/notifymesg"
	"errors"
	"fmt"

	"log"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func Managebg(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("managing background task stopped")
			return
		case <-ticker.C:
			var count int64
			if err := models.Db.Model(&models.Pickuprequest{}).Where("status = ?", false).Count(&count).Error; err != nil {
				log.Println("Error checking pending requests:", err)
				continue
			}
			if count > 0 {
				log.Printf("Starting background tasks for %d pending requests", count)

				var wg sync.WaitGroup
				wg.Add(3)

				go func() {
					defer wg.Done()
					Getlocations(ctx)
				}()

				go func() {
					defer wg.Done()
					Autocomplete(ctx)
				}()

				go func() {
					defer wg.Done()
					Decidequality(ctx)
				}()

				wg.Wait() // Ensure all tasks complete
			} else {
				log.Println("No pending requests, background tasks are idle")
			}
		}
	}
}

func Autocomplete(ctx context.Context) {
	const maxconcurrentrequests = 5
	semaphore := make(chan struct{}, maxconcurrentrequests)

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping autocomplete goroutine.")
			return
		default:
			var nextrequest models.Pickuprequest

			err := models.Db.Where("status = ?", false).Order("priority asc").First(&nextrequest).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println("No pending request... queue is empty")
				time.Sleep(5 * time.Minute)
				continue
			} else if err != nil {
				log.Printf("Error fetching next request: %v\n", err)
				time.Sleep(1 * time.Minute)
				continue
			}
			if nextrequest.TravelTime == 0.0 {
				log.Printf("request is still pending wait until time calculation id:%d", nextrequest.Id)
				time.Sleep(1 * time.Minute)
				continue
			}
			if len(semaphore) == cap(semaphore) {
				log.Printf("semaphore is currently full the next request will start as it get a slot.... upcoming request id:%d\n", nextrequest.Id)
			}
			semaphore <- struct{}{}
			go func(req models.Pickuprequest) {
				defer func() { <-semaphore }()

				timetocollect := time.Duration(nextrequest.TravelTime * float64(time.Second))
				log.Println("collection started... required time:", timetocollect)
				time.Sleep(timetocollect)
				//send message via Addnotification
				notifymesg.AddNotification(nextrequest.Name, fmt.Sprintf("collection started... required time:%s", timetocollect))
				err := beeep.Notify("Evault", fmt.Sprintf("collection started... required time:%s", timetocollect), "")
				if err != nil {
					log.Printf("Failed to send notification: %v", err)
				}

				if err := models.Db.Save(&nextrequest).Update("status", true).Error; err != nil {
					log.Printf("Error completing request: %d:%v\n", nextrequest.Id, err)
				} else {
					var completedreq models.Pickuprequest
					log.Printf("Completed request: %v \n", nextrequest.Id)
					if err := models.Db.Where("id=? AND status=?", nextrequest.Id, true).First(&completedreq).Error; err != nil {
						log.Printf("error fetching request %v\n", nextrequest.Id)
						return
					}
					collected := models.Collectedrequests{
						Name:        nextrequest.Name,
						Latitude:    nextrequest.Latitude,
						Longitude:   nextrequest.Longitude,
						Wastetype:   nextrequest.Wastetype,
						Description: nextrequest.Description,
						Phone:       nextrequest.Phone,
						Quantity:    nextrequest.Quantity,
						Status:      true,
						Priority:    nextrequest.Priority,
						CompletedAt: time.Now(),
					}
					if err := models.Db.Create(&collected).Error; err != nil {
						log.Printf("error while inserting collection: %d %v\n", nextrequest.Id, err)
					}
					if err := models.Db.Delete(&nextrequest).Error; err != nil {
						log.Printf("error while deleting record %v \n", nextrequest.Id)
						return
					}

				}
			}(nextrequest)
			time.Sleep(1 * time.Minute)
		}
	}
}
