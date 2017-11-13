package apiclient

import (
	"context"
	"fmt"

	"github.com/yuhanfang/riot/constants/region"
)

type SummonerDTO struct {
	ProfileIconID int    // ID of the summoner icon associated with the summoner.
	Name          string //Summoner name.
	SummonerLevel int64  // Summoner level associated with the summoner.
	RevisionDate  int64  // Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: profile icon change, playing the tutorial or advanced tutorial, finishing a game, summoner name change
	ID            int64  // Summoner ID.
	AccountID     int64  //Account ID.
}

func (c *client) GetByAccountID(ctx context.Context, r region.Region, accountID int64) (*SummonerDTO, error) {
	var res SummonerDTO
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/summoner/v3/summoners/by-account", fmt.Sprintf("/%d", accountID), nil, &res)
	return &res, err
}

func (c *client) GetBySummonerName(ctx context.Context, r region.Region, name string) (*SummonerDTO, error) {
	var res SummonerDTO
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/summoner/v3/summoners/by-name", fmt.Sprintf("/%s", name), nil, &res)
	return &res, err
}

func (c *client) GetBySummonerID(ctx context.Context, r region.Region, summonerID int64) (*SummonerDTO, error) {
	var res SummonerDTO
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/summoner/v3/summoners", fmt.Sprintf("/%d", summonerID), nil, &res)
	return &res, err
}
