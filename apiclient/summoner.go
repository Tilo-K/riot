package apiclient

import (
	"context"
	"fmt"
	"github.com/Tilo-K/riot/constants/v5region"

	"github.com/Tilo-K/riot/constants/region"
)

type Summoner struct {
	ProfileIconID int    `json:"profileIconID",datastore:",noindex"` // ID of the summoner icon associated with the summoner.
	Name          string `json:"name",datastore:",noindex"`          // Summoner name.
	PUUID         string `json:"puuid",datastore:",noinex"`          // PUUID is the player universally unique identifier.
	SummonerLevel int64  `json:"summonerLevel",datastore:",noindex"` // Summoner level associated with the summoner.
	AccountID     string `json:"accountID",datastore:",noindex"`     // Encrypted account ID.
	ID            string `json:"id",datastore:",noindex"`            // Encrypted summoner ID.
	RevisionDate  int64  `json:"revisionDate",datastore:",noindex"`  // Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: profile icon change, playing the tutorial or advanced tutorial, finishing a game, summoner name change
}

type RiotAccount struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

func (c *client) GetByAccountID(ctx context.Context, r region.Region, accountID string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/summoner/v4/summoners/by-account", fmt.Sprintf("/%s", accountID), nil, &res)
	return &res, err
}

func (c *client) GetBySummonerName(ctx context.Context, r region.Region, name string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/summoner/v4/summoners/by-name", fmt.Sprintf("/%s", name), nil, &res)
	return &res, err
}

func (c *client) GetBySummonerPUUID(ctx context.Context, r region.Region, puuid string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/summoner/v4/summoners/by-puuid", fmt.Sprintf("/%s", puuid), nil, &res)
	return &res, err
}

func (c *client) GetBySummonerID(ctx context.Context, r region.Region, summonerID string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/summoner/v4/summoners", fmt.Sprintf("/%s", summonerID), nil, &res)
	return &res, err
}

func (c *client) GetRiotAccountByNameAndTag(ctx context.Context, r v5region.V5Region, name string, tag string) (*RiotAccount, error) {
	var res RiotAccount
	_, err := c.dispatchAndUnmarshalV5(ctx, r, "/riot/account/v1/accounts/by-riot-id", fmt.Sprintf("/%s/%s", name, tag), nil, &res)
	return &res, err
}

func (c *client) GetRiotAccountByPuuid(ctx context.Context, r v5region.V5Region, puuid string) (*RiotAccount, error) {
	var res RiotAccount
	_, err := c.dispatchAndUnmarshalV5(ctx, r, "/riot/account/v1/accounts/by-puuid", fmt.Sprintf("/%s", puuid), nil, &res)
	return &res, err
}
