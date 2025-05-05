package pkg

import (
	"common/model"
	"log"

	"github.com/robfig/cron/v3"
)

func StartWeeklyUpdate() {
	c := cron.New()
	_, err := c.AddFunc("0 0 0 * * SUN", func() {
		var commission model.Commission
		result := commission.CalculateAndRankTotalCommission()
		if len(result) == 0 {
			log.Printf("Failed to calculate and rank total commission: no result")
		}
	})
	if err != nil {
		return
	}
	c.Start()
}
