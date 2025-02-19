package mapview

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	serverMutex   sync.Mutex   // Protects the HTTP server startup logic
	serverStarted bool         // Flag to track if the server is running
	server        *http.Server // Store the server for graceful shutdown
)

func StartMapServer() {
	serverMutex.Lock()
	defer serverMutex.Unlock()

	if serverStarted {
		log.Println("Server is already running.")
		return
	}

	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
<!DOCTYPE html>
<html>
<head>
  <title>Interactive Map</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="https://unpkg.com/leaflet/dist/leaflet.css" />
  <script src="https://unpkg.com/leaflet/dist/leaflet.js"></script>
</head>
<body>
  <div id="map" style="width: 100%; height: 90vh;"></div>
  <button id="toggleModeBtn" style="position: absolute; top: 10px; right: 10px; z-index: 9999; padding: 10px 15px; font-size: 16px; background-color: #008CBA; color: white; border: none; border-radius: 5px;">Selection Mode</button>
  <script>
    const map = L.map('map').setView([20.5937, 78.9629], 5); 
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 18,
    }).addTo(map);

    let selectMode = true; 
    let marker = null;

    document.getElementById('toggleModeBtn').addEventListener('click', function() {
        selectMode = !selectMode;
        const modeText = selectMode ? 'Selection Mode' : 'Drag Mode';
        this.innerText = modeText;

        console.log("Mode changed to:", modeText);
    });

    map.on('click', function (e) {
        if (selectMode) {
            if (marker) {
                map.removeLayer(marker);
            }
            marker = L.marker(e.latlng).addTo(map);
            
            const coords = { lat: e.latlng.lat, lng: e.latlng.lng };
            console.log("Coordinates selected:", JSON.stringify(coords));

            fetch('http://localhost:8082/setCoordinates', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(coords),
                credentials: 'include',
            }).then(response => response.json())
              .then(data => console.log("Coordinates sent to Go:", data.message))
              .catch(error => console.error("Error sending coordinates:", error));
        }
    });
  </script>
</body>
</html>
	`)
	})

	server = &http.Server{Addr: ":8081"}

	go func() {
		log.Println("Starting map server on port 8081...")
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Error starting the map server: %v", err)
		}
	}()

	serverStarted = true
	log.Println("Map server started on port 8081.")
}

func StopMapServer() {
	if server != nil {
		log.Println("Stopping the map server...")
		if err := server.Close(); err != nil {
			log.Printf("Error stopping the server: %v", err)
		} else {
			log.Println("Map server stopped.")
		}
	}
}

func ListenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	StopMapServer()
}
