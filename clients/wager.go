package clients

import (
	"context"
	"errors"
	"github.com/ryannguyen89/wager-service/clients/wager/client"
	"sync"
)

var (
	wagerClient *WagerClient
	wagerOnce sync.Once
	wagerCfg WagerConfig
)

type WagerConfig struct {
	Address string
}

type WagerClient struct {
	*wager.APIClient
}

func SetWagerConfig(cfg WagerConfig) {
	wagerCfg = cfg
}

func GetWagerClient() *WagerClient {
	wagerOnce.Do(func() {
		wagerClient = initializeWagerClient()
	})
	return wagerClient
}

func initializeWagerClient() *WagerClient {
	var (
		cfg = wager.NewConfiguration()
	)
	cfg.BasePath = wagerCfg.Address

	apiCli := wager.NewAPIClient(cfg)
	return &WagerClient{APIClient: apiCli}
}

func (w *WagerClient) CreateWager(ctx context.Context, req *wager.CreateWagerRequest) (*wager.Wager, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}
	resp, _, err := w.WagerApi.CreateWagerPost(ctx, *req)
	return &resp, err
}
