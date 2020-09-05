package clients

import (
	"context"
	wager "github.com/ryannguyen89/wager-service/clients/wager/client"
	"testing"
)

func init() {
	SetWagerConfig(WagerConfig{Address: "http://localhost:8080"})
}

func TestWagerClient_CreateWager(t *testing.T) {
	type args struct {
		ctx context.Context
		req *wager.CreateWagerRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *wager.Wager
		wantErr bool
	}{
		{
			"nil request",
			args{
				ctx: context.Background(),
				req: nil,
			},
			nil,
			true,
		},
		{
			"valid request",
			args{
				ctx: context.Background(),
				req: &wager.CreateWagerRequest{
					TotalWagerValue: 1000,
					Odds: 100,
					SellingPercentage: 50,
					SellingPrice: 550,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := GetWagerClient()
			_, err := w.CreateWager(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWager() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}