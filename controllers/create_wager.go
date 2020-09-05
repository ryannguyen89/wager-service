package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryannguyen89/wager-service/models"
	"github.com/ryannguyen89/wager-service/models/errors"
	"net/http"
)

// CreateWagerPost -
func CreateWagerPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
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
