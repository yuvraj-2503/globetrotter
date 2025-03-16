package probe

import (
	"context"
	"encoding/json"
	"globetrotter/game/models"
	"globetrotter/game/service"
	"io/ioutil"
	"log"
	"os"
)

type DestinationsProbe struct {
	destinationService service.DestinationService
}

func NewDestinationsProbe(destinationService service.DestinationService) *DestinationsProbe {
	return &DestinationsProbe{destinationService: destinationService}
}

func (probe *DestinationsProbe) FetchDestinationsFromFile(ctx *context.Context, jsonFilePath string) error {
	file, err := os.Open(jsonFilePath)
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
		return err
	}

	var destinations []*models.Destination
	if err := json.Unmarshal(data, &destinations); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
		return err
	}

	err = probe.destinationService.InsertBulk(ctx, destinations)
	if err != nil {
		return err
	}
	return nil
}
