package dashboard

import (
	"log"

	"github.com/Whitfrost21/EVault/evault/models"
)

func FetchBarheight(labels []string) []float32 {
	var history []models.History
	if res := models.Db.Find(&history); res.Error != nil {
		log.Println("error while opening database", res.Error)
		return nil
	}
	counts := make([]float32, len(labels))

	// Iterate over each history record to count occurrences of each Wastetype
	for _, record := range history {
		// Iterate over label1 to match Wastetype and increment the count
		for i, wasteType := range labels {
			if record.Wastetype == wasteType {
				counts[i]++ // Increment the count at the corresponding index
				break       // No need to check further once a match is found
			}
		}
	}

	return counts
}
