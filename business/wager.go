package business

import (
	"context"
	"github.com/astaxie/beego/orm"
	"github.com/ryannguyen89/wager-service/models"
	"github.com/ryannguyen89/wager-service/models/errors"
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
	id, err := wager.Create(ctx)
	if err != nil {
		log.Printf("Create wager to Db error: %v", err)
		message = err.Error()
		return
	}
	wager.Id = id
	code = http.StatusCreated

	log.Printf("Wager %s", utils.JsonSerialize(wager))
	return
}

func BuyWager(ctx context.Context, wagerId int64, buyingPrice float32) (wPurchase *models.WagerPurchase, code int, message string) {
	var (
		o =orm.NewOrm()
		w = models.Wager{
			Id: wagerId,
		}
	)
	code = http.StatusInternalServerError

	// Begin a transaction
	if err := o.BeginTx(ctx, nil); err != nil {
		message = err.Error()
		return
	}

	// Read wager
	err := w.ReadForUpdate(ctx)
	if err != nil {
		code = http.StatusBadRequest
		message = err.Error()
		_ = o.Rollback()
		return
	}
	log.Printf("wager %s", utils.JsonSerialize(w))

	// Check if buyingPrice is valid
	if buyingPrice > w.CurrentSellingPrice {
		code = http.StatusNotAcceptable
		message = errors.BuyingPriceNotAccept

		_ = o.Rollback()
		return
	}

	// Check if valid to sell
	if buyingPrice + w.AmountSold > float32(w.TotalWagerValue) {
		code = http.StatusNotAcceptable
		message = errors.BuyingPriceNotAccept

		_ = o.Rollback()
		return
	}

	// Create wager purchase object
	wPurchase = &models.WagerPurchase{
		WagerId:     wagerId,
		BuyingPrice: buyingPrice,
		BoughtAt:    time.Now().Unix(),
	}
	id, err := wPurchase.Create(ctx)
	if err != nil {
		log.Printf("Create wager purchase error: %v", err)
		message = err.Error()

		_ = o.Rollback()
		return
	}
	wPurchase.Id = id
	log.Printf("wPurchase %s", utils.JsonSerialize(wPurchase))

	// Update wager
	w.CurrentSellingPrice = buyingPrice
	w.AmountSold += buyingPrice
	w.PercentageSold += int32(buyingPrice / float32(w.TotalWagerValue) * 100)
	err = w.Update(ctx)
	if err != nil {
		log.Printf("update wager error: %v", err)
		message = err.Error()

		_ = o.Rollback()
		return
	}

	code = http.StatusCreated
	_ = o.Commit()
	return
}

func ListWagers(ctx context.Context, req *models.ListWagerRequest) (lWagers []*models.Wager, code int, message string) {
	var (
		err error
	)
	code = http.StatusInternalServerError

	lWagers, err = models.ListWagers(ctx, req.Page, req.Limit)
	if err != nil {
		log.Printf("List wagers error: %v", err)
		message = err.Error()
		return
	}

	code = http.StatusOK
	return
}
