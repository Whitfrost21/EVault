package models

import (
	"time"

	"gorm.io/gorm"
)

type Locations struct {
}
type Pickuprequest struct {
	Id          uint    `gorm:"primaryKey"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Wastetype   string  `json:"wastetype"`
	Description string  `json:"description"`
	Phone       string  `gorm:"size:10;not null" json:"phone"`
	Quantity    int     `gorm:"not null" json:"quantity"`
	Status      bool    `json:"status"`
	Priority    int     `json:"priority"`
	Distance    float64 `json:"distance"`
	TravelTime  float64 `json:"traveltime"`
}

var Db *gorm.DB

type Wastetype struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type Collectedrequests struct {
	Id          uint      `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Address     string    `json:"address"`
	Wastetype   string    `json:"wastetype"`
	Description string    `json:"description"`
	Phone       string    `gorm:"size:10;not null" json:"phone"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	Priority    int       `json:"priority"`
	Weight      float64   `json:"weight"`
	Status      bool      `json:"status"`
	Quality     string    `json:"quality"`
	Cost        float64   `gorm:"default:0.0"`
	CompletedAt time.Time `gorm:"autoCreateTime" json:"completedat"`
}

type History struct {
	Id          uint      `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Address     string    `json:"address"`
	Wastetype   string    `json:"wastetype"`
	Description string    `json:"description"`
	Phone       string    `gorm:"size:10;not null" json:"phone"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	Cost        float64   `gorm:"default:0.0"`
	Quality     string    `json:"quality"`
	Weight      float64   `json:"weight"`
	CompletedAt time.Time `gorm:"autoCreateTime" json:"completedat"`
}

type Graphopresponse struct {
	Paths []struct {
		Distance float64 `json:"distance"`
		Time     float64 `json:"time"`
	} `json:"paths"`
}
