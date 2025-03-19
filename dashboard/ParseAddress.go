package dashboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NominatimResponse struct {
	DisplayName string `json:"display_name"`
}

func GetAddress(lat, long float64) (string, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%f&lon=%f&format=json", lat, long)
	fmt.Println("Requesting:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Add a User-Agent header
	req.Header.Add("User-Agent", "EVault/1.0 (pzgaming62@gmail.com)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the response is OK
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error: Received non-OK HTTP status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Print the raw response to check for potential issues
	// fmt.Println("Response Body: ", string(body))

	var result NominatimResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	// Check if display_name is available
	if result.DisplayName == "" {
		return "", fmt.Errorf("no address found for coordinates %f, %f", lat, long)
	}

	return result.DisplayName, nil
}
