package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryannguyen89/wager-service/business"
	"github.com/ryannguyen89/wager-service/models"
	"github.com/ryannguyen89/wager-service/models/errors"
	"github.com/ryannguyen89/wager-service/utils"
	"log"
	"net/http"
)

type idRequest struct {
	WagerId int64 `uri:"wager_id" binding:"required"`
}

// BuyWagersPost -
func BuyWagersPost(c *gin.Context) {
	var (
		req models.BuyWagerRequest
		idReq idRequest
	)

	if err := c.ShouldBindUri(&idReq); err != nil {
		log.Printf("Bind request wager id error: %v", err)
		c.JSON(http.StatusBadRequest, models.Error{Error:err.Error()})
		return
	}
	log.Printf("wager_id %d", idReq.WagerId)

	if err := c.ShouldBind(&req); err != nil {
		log.Printf("Bind request error: %v", err)
		c.JSON(http.StatusBadRequest, models.Error{Error:err.Error()})
		return
	}
	log.Printf("Request %v", utils.JsonSerialize(req))

	wPurchase, code, msg := business.BuyWager(c.Request.Context(), idReq.WagerId, req.BuyingPrice)
	if code != http.StatusCreated {
		c.JSON(code, models.Error{Error:msg})
		return
	}
	c.JSON(code, wPurchase)
}

func validateBuyWagerRequest(r *models.BuyWagerRequest) (code int, message string) {
	buyingPriceMul100 := r.BuyingPrice * 100
	if buyingPriceMul100- float32(int(buyingPriceMul100)) > 1e-5 {
		return http.StatusBadRequest, errors.InvalidBuyingPriceField
	}
	return
}
