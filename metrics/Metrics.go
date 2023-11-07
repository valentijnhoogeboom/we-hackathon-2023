package metrics

import (
	"GlobalAPI/database"
	"GlobalAPI/models"
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var GlobalMetrics *Metrics

type Metrics struct {
	Power api.WriteAPIBlocking
}

func (metrics *Metrics) Start() {
	metrics.Power.Flush(context.Background())
	var meters []models.Meter

	database.StaticDatabase.DB.Model(&models.Meter{}).Find(&meters)

	for _, meter := range meters {
		var data []models.Data
		database.StaticDatabase.DB.Model(&models.Data{}).Find(&data, &models.Data{MeterID: meter.ID})

		for _, datum := range data {
			p := influxdb2.NewPoint(fmt.Sprintf("%v", datum.MeterID),
				map[string]string{"unit": "watt"},
				map[string]interface{}{"current": (datum.ActivePowerW) * (10000.0) / 3600.0},
				datum.CreatedAt)
			err := metrics.Power.WritePoint(context.Background(), p)
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
		}

	}
}

func (metrics *Metrics) WritePoint(ctx context.Context, point ...*write.Point) error {
	defer func() {
		recover()
	}()

	return metrics.Power.WritePoint(ctx, point...)
}
