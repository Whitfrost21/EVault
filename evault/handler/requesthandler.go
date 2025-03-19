package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Whitfrost21/EVault/evault/models"
	"github.com/gin-gonic/gin"
)

func Getallrequest(c *gin.Context) {
	var allrequests []models.Pickuprequest
	if result := models.Db.Find(&allrequests); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, allrequests)
}

func Createrequest(c *gin.Context) {
	var request models.Pickuprequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := map[string]bool{
		"address cannot be empty": request.Latitude == 0.0 && request.Longitude == 0.0,
		"name cannot be empty":    request.Name == "",
		"wastetype cannot be empty check wastetypes on /getwastetypes": request.Wastetype == "",
		"other wastetype must have some description":                   request.Wastetype == "other" && request.Description == "",
		"quantity cannot be 0":                                         request.Quantity == 0,
		"phone number must be 10 digits":                               len(request.Phone) != 10,
	}
	for err, isvalid := range validate {
		if isvalid {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}
	request.Status = false
	switch request.Wastetype {
	case "Batteries":
		request.Priority = 1
	case "Electronics and sensors":
		request.Priority = 2
	case "Electric motors":
		request.Priority = 3
	default:
		request.Priority = 10
	}
	if result := models.Db.Create(&request); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "pickup added", "data": request})
}
func Getrequest(c *gin.Context) {
	var request models.Pickuprequest
	id := c.Param("id")
	id = strings.TrimSpace(id)
	index, err := strconv.Atoi(id)
	if err != nil {
		panic(err)

	}
	if err := models.Db.First(&request, "id=?", index).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, request)
}

func Updaterequest(c *gin.Context) {
	var request models.Pickuprequest
	id := c.Param("id")
	id = strings.TrimSpace(id)
	index, err := strconv.Atoi(id)
	if err != nil {
		panic(err)

	}

	var update struct {
		Name        string  `json:"name"`
		Address     string  `json:"address"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		Wastetype   string  `json:"wastetype"`
		Description string  `json:"description"`
		Phone       int     `gorm:"size:10;not null" json:"phone"`
		Quantity    int     `gorm:"not null" json:"quantity"`
		Status      bool    `json:"status"`
	}

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	models.Db.Model(&models.Pickuprequest{}).Where("id=?", index).Updates(update)
	c.JSON(http.StatusOK, gin.H{"updated": request})
}

func Deleterequest(c *gin.Context) {
	var request models.Pickuprequest
	id := c.Param("id")
	id = strings.TrimSpace(id)
	index, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	if err := models.Db.First(&request, "id=?", index).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no request found here"})
		return
	}
	if err := models.Db.Delete(&request).Error; err != nil {
		log.Printf("error while deleting %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully!"})
}

func Getcollected(c *gin.Context) {
	var collectedlist []models.Collectedrequests
	if res := models.Db.Find(&collectedlist); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
		return
	}
	c.JSON(http.StatusOK, collectedlist)
}

func Gethistory(c *gin.Context) {
	var history []models.History
	if res := models.Db.Find(&history); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
		return
	}
	c.JSON(http.StatusOK, history)
}

func GetTotalweight(c *gin.Context) {
	var list []models.History
	var TotalWeight float64
	if res := models.Db.Find(&list); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
		return
	}
	TotalWeight = 0
	for _, req := range list {
		TotalWeight += req.Weight
	}
	c.JSON(http.StatusOK, TotalWeight)
}

func Getwastetypes(c *gin.Context) {
	var wastetypes = []models.Wastetype{
		{Name: "Batteries", Description: "Used or defective EV batteries", Category: "High Value"},
		{Name: "Electric Motors", Description: "Defunct motors from EVs", Category: "Mechanical"},
		{Name: "Charging Accessories", Description: "Cables, adapters, and chargers", Category: "Accessories"},
		{Name: "Wiring and Harnesses", Description: "Electrical wirings and connectors", Category: "Electrical"},
		{Name: "Electronics and Sensors", Description: "Circuit boards, sensors, and controllers", Category: "Electronics"},
		{Name: "HVAC Systems", Description: "Heating, ventilation, and cooling parts", Category: "HVAC"},
		{Name: "Other", Description: "Any other e-waste not listed above", Category: "Miscellaneous"},
	}
	c.JSON(http.StatusOK, wastetypes)
}
