package controllers

import (
	"github.com/ryannguyen89/wager-service/models"
	"github.com/ryannguyen89/wager-service/models/errors"
	"net/http"
	"testing"
)

func Test_validateBuyWagerRequest(t *testing.T) {
	type args struct {
		r *models.BuyWagerRequest
	}
	tests := []struct {
		name        string
		args        args
		wantCode    int
		wantMessage string
	}{
		{"selling_price more than two decimal places",
			args{r: &models.BuyWagerRequest{BuyingPrice: 37.123}},
			http.StatusBadRequest,
			errors.InvalidBuyingPriceField,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotMessage := validateBuyWagerRequest(tt.args.r)
			if gotCode != tt.wantCode {
				t.Errorf("validateBuyWagerRequest() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			if gotMessage != tt.wantMessage {
				t.Errorf("validateBuyWagerRequest() gotMessage = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
