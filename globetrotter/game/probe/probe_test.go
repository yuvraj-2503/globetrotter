package probe

import (
	"context"
	"globetrotter/game/db"
	"globetrotter/game/service"
	mongodb "mongo-utils"
	"testing"
)

var destinationService service.DestinationService

var ctx context.Context

func init() {
	mongoConfig := mongodb.MongoConfig{
		ConnectionString: "mongodb://localhost:27017",
		Database:         "globetrotter",
		Username:         "",
		Password:         "",
	}
	coll, err := mongoConfig.GetCollection("destinations")
	if err != nil {
		panic(err)
	}
	ctx = context.Background()
	destDB := db.NewDestinationDBStore(coll)
	destinationService = service.NewDestinationService(destDB)
}

func TestDestinationsProbe_FetchDestinationsFromFile(t *testing.T) {
	type fields struct {
		destinationService service.DestinationService
	}
	type args struct {
		ctx          *context.Context
		jsonFilePath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				destinationService: destinationService,
			},
			args: args{
				ctx:          &ctx,
				jsonFilePath: "destinations.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			probe := &DestinationsProbe{
				destinationService: tt.fields.destinationService,
			}
			if err := probe.FetchDestinationsFromFile(tt.args.ctx, tt.args.jsonFilePath); (err != nil) != tt.wantErr {
				t.Errorf("FetchDestinationsFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
