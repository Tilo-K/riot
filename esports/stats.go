package esports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Tilo-K/riot/apiclient"
	"github.com/Tilo-K/riot/constants/region"
)

func (c Client) GetGameStats(ctx context.Context, region region.Region, gameID int64, gameHash string) (*apiclient.Match, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://acs.leagueoflegends.com/v1/stats/game/%s/%d?gameHash=%s", region, gameID, gameHash), nil)
	if err != nil {
		return nil, err
	}
	var res apiclient.Match
	_, err = c.doJSON(ctx, req, &res)
	return &res, err
}

func (c Client) GetGameTimeline(ctx context.Context, region region.Region, gameID int64, gameHash string) (*apiclient.MatchTimeline, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://acs.leagueoflegends.com/v1/stats/game/%s/%d/timeline?gameHash=%s", region, gameID, gameHash), nil)
	if err != nil {
		return nil, err
	}
	var res apiclient.MatchTimeline
	_, err = c.doJSON(ctx, req, &res)
	return &res, err
}
