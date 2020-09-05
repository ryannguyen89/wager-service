package business

import (
	"context"
	"github.com/ryannguyen89/wager-service/models"
	"github.com/ryannguyen89/wager-service/utils"
	"log"
	"net/http"
	"time"
)

func CreateWager(ctx context.Context, r *models.CreateWagerRequest) (wager *models.Wager, code int, message string) {
	code = http.StatusInternalServerError

	wager = &models.Wager{
		TotalWagerValue:     r.TotalWagerValue,
		Odds:                r.Odds,
		SellingPercentage:   r.SellingPercentage,
		SellingPrice:        r.SellingPrice,
		CurrentSellingPrice: r.SellingPrice,
		PlacedAt:            time.Now().Unix(),
	}
	id, err := wager.Insert(ctx)
	if err != nil {
		log.Printf("Insert wager to Db error: %v", err)
		message = err.Error()
		return
	}
	wager.Id = id
	code = http.StatusCreated

	log.Printf("Wager %s", utils.JsonSerialize(wager))
	return
}
