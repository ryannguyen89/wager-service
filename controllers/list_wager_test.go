package controllers

import (
	"github.com/ryannguyen89/wager-service/models"
	"github.com/ryannguyen89/wager-service/models/errors"
	"net/http"
	"testing"
)

func Test_validateListWagerRequest(t *testing.T) {
	type args struct {
		r *models.ListWagerRequest
	}
	tests := []struct {
		name        string
		args        args
		wantCode    int
		wantMessage string
	}{
		{"negative page",
			args{r: &models.ListWagerRequest{Page: -10, Limit: 100}},
			http.StatusBadRequest,
			errors.InvalidPage,
		},
		{"zero page",
			args{r: &models.ListWagerRequest{Page: 0, Limit: 100}},
			http.StatusBadRequest,
			errors.InvalidPage,
		},
		{"negative limit",
			args{r: &models.ListWagerRequest{Page: 1, Limit: -10}},
			http.StatusBadRequest,
			errors.InvalidLimit,
		},
		{"zero limit",
			args{r: &models.ListWagerRequest{Page: 1, Limit: 0}},
			http.StatusBadRequest,
			errors.InvalidLimit,
		},
		{"valid request",
			args{r: &models.ListWagerRequest{Page: 1, Limit: 10}},
			http.StatusOK,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotMessage := validateListWagerRequest(tt.args.r)
			if gotCode != tt.wantCode {
				t.Errorf("validateListWagerRequest() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			if gotMessage != tt.wantMessage {
				t.Errorf("validateListWagerRequest() gotMessage = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
