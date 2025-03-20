package mapview

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gen2brain/beeep"
	"github.com/rs/cors"
)

var (
	Lat, Lng       float64
	coordServerMux *http.ServeMux
	coordServer    *http.Server
	started        bool
	servermutex    sync.Mutex
)

type NominatimResponse struct {
	DisplayName string `json:"display_name"`
}

// StartCoordinateServer starts the coordinate server and listens for incoming requests
func StartCoordinateServer() {
	servermutex.Lock()
	defer servermutex.Unlock()

	if started {
		log.Println("Coordinate server is already running.")
		return
	}

	// Initialize the mux and define the /setCoordinates handler
	coordServerMux = http.NewServeMux()
	coordServerMux.HandleFunc("/setCoordinates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var coords struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			}

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&coords)
			if err != nil {
				http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
				return
			}

			Lat = coords.Lat
			Lng = coords.Lng
			log.Printf("Received coords: lat=%.6f, lon=%.6f", Lat, Lng)
			address := searchaddress(Lat, Lng)
			err = beeep.Notify("EVault", fmt.Sprintf("Selected Location: %s", address), "")
			if err != nil {
				log.Println("error while sending beeep notification", err)
			}
			response := map[string]string{
				"status":  "success",
				"message": fmt.Sprintf("Coordinates received: Lat=%.6f, Lng=%.6f", Lat, Lng),
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// Set up CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"}, // Allow requests from this origin
		AllowedMethods:   []string{"POST", "OPTIONS"},       // Explicitly allow POST and OPTIONS methods
		AllowedHeaders:   []string{"Content-Type"},          // Allow Content-Type header
		AllowCredentials: true,                              // Allow credentials (e.g., cookies)
	})

	handler := c.Handler(coordServerMux)

	// Start the HTTP server in a goroutine
	coordServer = &http.Server{Addr: ":8082", Handler: handler}

	go func() {
		log.Println("Starting coordinate server on port 8082 with CORS enabled...")
		if err := coordServer.ListenAndServe(); err != nil {
			log.Fatalf("Error starting coordinate server: %v", err)
		}
	}()

	started = true
	log.Println("Coordinate server started on port 8082.")
}

// Gracefully shutdown the coordinate server
func StopCoordinateServer() {
	if coordServer != nil {
		log.Println("Stopping coordinate server...")
		if err := coordServer.Close(); err != nil {
			log.Printf("Error stopping coordinate server: %v", err)
		} else {
			log.Println("Coordinate server stopped.")
		}
	}
}

func searchaddress(lat, long float64) string {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%f&lon=%f&format=json", lat, long)
	fmt.Println("Requesting:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error while searching address", err)
	}

	// Add a User-Agent header
	req.Header.Add("User-Agent", "EVault/1.0 (your email)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error while searching address", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("error while searching address", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error while searching address", err)
	}

	var result NominatimResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("error while marshalling json")
	}

	if result.DisplayName == "" {
		log.Println("error no address on this coordinates")
	}

	return result.DisplayName
}

// ListenForShutdown listens for shutdown signals to gracefully stop the server
func Listencoords() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown signal received. Stopping servers...")

	// Stop both the map and coordinate servers gracefully
	StopMapServer()
	StopCoordinateServer()

	log.Println("Application has been gracefully stopped.")
}
