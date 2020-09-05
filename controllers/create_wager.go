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

// CreateWagerPost -
func CreateWagerPost(c *gin.Context) {
	var (
		req models.CreateWagerRequest
	)

	if err := c.ShouldBind(&req); err != nil {
		log.Printf("Bind request error: %v", err)
		c.JSON(http.StatusBadRequest, models.Error{Error:"Bind request error"})
		return
	}
	log.Printf("Request %v", utils.JsonSerialize(req))

	// Validate request
	code, msg := validateCreateWagerRequest(&req)
	if code != http.StatusOK {
		log.Printf("Validate request error: %s", msg)
		c.JSON(code, models.Error{Error:msg})
		return
	}

	wager, code, msg := business.CreateWager(c.Request.Context(), &req)
	if code != http.StatusCreated {
		c.JSON(code, models.Error{Error:msg})
		return
	}
	c.JSON(code, wager)
}

func validateCreateWagerRequest(r *models.CreateWagerRequest) (code int, message string) {
	if r.TotalWagerValue <= 0 {
		return http.StatusBadRequest, errors.InvalidTotalWagerValueField
	}

	if r.Odds <= 0 {
		return http.StatusBadRequest, errors.InvalidOddsField
	}

	if r.SellingPercentage < 1 || r.SellingPercentage > 100 {
		return http.StatusBadRequest, errors.InvalidSellingPercentageField
	}

	sellingPriceMul100 := r.SellingPrice * 100
	if sellingPriceMul100 - float32(int(sellingPriceMul100)) > 1e-5 {
		return http.StatusBadRequest, errors.InvalidSellingPriceField
	}

	sellingAmount := float32(r.TotalWagerValue) * float32(r.SellingPercentage) / 100
	if r.SellingPrice <= sellingAmount {
		return http.StatusBadRequest, errors.InvalidSellingPriceField
	}

	code = http.StatusOK
	return
}
