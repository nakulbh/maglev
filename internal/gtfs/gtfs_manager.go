package gtfs

import (
	"context"
	"fmt"
	"github.com/jamespfennell/gtfs"
	"io"
	"log"
	"net/http"
	"time"
)

// Manager manages the GTFS data and provides methods to access it
type Manager struct {
	gtfsURL     string
	gtfsData    *gtfs.Static
	lastUpdated time.Time
}

// InitGTFSManager initializes the Manager with the GTFS data from the given URL
func InitGTFSManager(gtfsURL string) (*Manager, error) {
	staticData, err := downloadAndParseStaticGTFS(gtfsURL)

	if err != nil {
		return nil, err
	}

	return &Manager{
		gtfsURL:     gtfsURL,
		gtfsData:    staticData,
		lastUpdated: time.Now(),
	}, nil
}

// downloadAndParseStaticGTFS downloads and parses the GTFS data from the given gtfsURL
func downloadAndParseStaticGTFS(gtfsURL string) (*gtfs.Static, error) {
	resp, _ := http.Get(gtfsURL)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	staticData, err := gtfs.ParseStatic(b, gtfs.ParseStaticOptions{})
	if err != nil {
		return nil, err
	}

	return staticData, nil
}

// updateGTFSPeriodically updates the GTFS data on a regular schedule
func (manager *Manager) updateGTFSPeriodically() {
	// Update every 24 hours
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Create a context with timeout for the download
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)

			// Download and parse the GTFS feed
			staticData, err := downloadAndParseStaticGTFS(manager.gtfsURL)
			cancel() // Always cancel the context when done

			if err != nil {
				// Log error but don't crash the application
				log.Printf("Error updating GTFS data: %v", err)
				continue
			}

			// Update the GTFS data in the manager
			manager.gtfsData = staticData
			manager.lastUpdated = time.Now()

			log.Printf("GTFS data updated successfully for %v", manager.gtfsURL)
		}
	}
}

func (manager *Manager) PrintStatistics() {
	fmt.Printf("URL: %s\n", manager.gtfsURL)
	fmt.Printf("Last Updated: %s\n", manager.lastUpdated)
	fmt.Println("Stops Count: ", len(manager.gtfsData.Stops))
	fmt.Println("Routes Count: ", len(manager.gtfsData.Routes))
	fmt.Println("Trips Count: ", len(manager.gtfsData.Trips))
	fmt.Println("Agencies Count: ", len(manager.gtfsData.Agencies))
}
