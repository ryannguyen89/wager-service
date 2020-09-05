package controllers

import (
	"github.com/ryannguyen89/wager-service/models"
	"github.com/ryannguyen89/wager-service/models/errors"
	"net/http"
	"testing"
)

func Test_validateCreateWagerRequest(t *testing.T) {
	type args struct {
		r *models.CreateWagerRequest
	}
	tests := []struct {
		name        string
		args        args
		wantCode    int
		wantMessage string
	}{
		{"negative total_wager_value",
			args{r: &models.CreateWagerRequest{TotalWagerValue: -10, Odds: 100}},
			http.StatusBadRequest,
			errors.InvalidTotalWagerValueField,
		},
		{"zero total_wager_value",
			args{r: &models.CreateWagerRequest{TotalWagerValue: 0}},
			http.StatusBadRequest,
			errors.InvalidTotalWagerValueField,
		},

		{"negative odds",
			args{r: &models.CreateWagerRequest{TotalWagerValue: 1000, Odds: -5}},
			http.StatusBadRequest,
			errors.InvalidOddsField,
		},
		{"zero odds",
			args{r: &models.CreateWagerRequest{TotalWagerValue: 1000, Odds: 0}},
			http.StatusBadRequest,
			errors.InvalidOddsField,
		},
		{"negative selling_percentage",
			args{r: &models.CreateWagerRequest{TotalWagerValue: 1000, Odds: 100, SellingPercentage: -7}},
			http.StatusBadRequest,
			errors.InvalidSellingPercentageField,
		},
		{"zero selling_percentage",
			args{r: &models.CreateWagerRequest{TotalWagerValue: 1000, Odds: 100, SellingPercentage: 0}},
			http.StatusBadRequest,
			errors.InvalidSellingPercentageField,
		},
		{"above 100 selling_percentage",
			args{r: &models.CreateWagerRequest{TotalWagerValue: 1000, Odds: 100, SellingPercentage: 101}},
			http.StatusBadRequest,
			errors.InvalidSellingPercentageField,
		},
		{"selling_price more than two decimal places",
			args{r: &models.CreateWagerRequest{TotalWagerValue: 1000, Odds: 100, SellingPercentage: 50,
				SellingPrice: 37.123}},
			http.StatusBadRequest,
			errors.InvalidSellingPriceField,
		},
		{"selling_price is invalid",
			args{r: &models.CreateWagerRequest{TotalWagerValue: 1000, Odds: 100, SellingPercentage: 50,
				SellingPrice: 50}},
			http.StatusBadRequest,
			errors.InvalidSellingPriceField,
		},
		{"valid request",
			args{r: &models.CreateWagerRequest{TotalWagerValue: 1000, Odds: 100, SellingPercentage: 50,
				SellingPrice: 550}},
			http.StatusOK,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotMessage := validateCreateWagerRequest(tt.args.r)
			if gotCode != tt.wantCode {
				t.Errorf("validateCreateWagerRequest() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			if gotMessage != tt.wantMessage {
				t.Errorf("validateCreateWagerRequest() gotMessage = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
