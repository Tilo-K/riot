package esports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Tilo-K/riot/constants/language"
	"github.com/Tilo-K/riot/constants/region"
	"github.com/Tilo-K/riot/esports/league"
	"github.com/Tilo-K/riot/types"
)

type Leagues struct {
	Leagues               []Leagues_League
	HighlanderTournaments []Leagues_HighlanderTournament
	HighlanderRecords     []Leagues_HighlanderRecord
	Teams                 []Leagues_Team
	// Players
}

type Leagues_League struct {
	ID          int64
	Slug        string
	Name        string
	GUID        string
	Region      string
	DrupalID    int64
	LogoURL     string
	CreatedAt   string // 2015-04-28T21:08:42.000Z
	UpdatedAt   string
	Abouts      map[language.Language]string
	Names       map[language.Language]string
	Tournaments []string
}

type Leagues_HighlanderTournament struct {
	ID              string
	Title           string
	Description     string
	LeagueReference string
	// RosteringStrategy
	// Queues
	// Rosters
	Published bool
	// Breakpoints
	Brackets map[string]Leagues_HighlanderTournament_Bracket
	// LiveMatches
	StartDate   string // YYYY-MM-DD
	EndDate     string // YYYY-MM-DD
	LeagueID    string
	PlatformIDs []string
	GameIDs     []string
	League      string
}

type Leagues_HighlanderTournament_Bracket struct {
	ID             string
	Name           string
	Position       int
	GroupPosition  int
	GroupName      string
	CanManufacture bool
	State          string
	// BracketType
	// MatchType
	// GameMode
	// Input
	Matches map[string]Leagues_HighlanderTournament_Bracket_Match
	// Standings
}

type Leagues_HighlanderTournament_Bracket_Match struct {
	ID            string
	Name          string
	Position      int
	State         string
	GroupPosition int
	Games         map[string]Leagues_HighlanderTournament_Bracket_Match_Game
}

type Leagues_HighlanderTournament_Bracket_Match_Game struct {
	ID            string
	Name          string
	GeneratedName string
	// GameMode
	// Input
	// Standings
	// Scores
	GameID     types.StringInt64
	GameRealm  region.Region
	PlatformID string
	Revision   int
}

type Leagues_HighlanderRecord struct {
	Wins       int
	Losses     int
	Ties       int
	Score      int
	Roster     string
	Tournament string
	Bracket    string
	ID         string
}

type Leagues_Team struct {
	ID           int64
	Slug         string
	Name         string
	GUID         string
	TeamPhotoURL string
	LogoURL      string
	Acroynm      string
	HomeLeague   string
	AltLogoURL   string
	CreatedAt    string // YYYY-MM-DDT18:34:47.000Z
	UpdatedAt    string
	Bios         map[language.Language]string
	// ForeignIDs
	Players  []int64
	Starters []int64
	Subs     []int64
}

func (c Client) GetLeagues(ctx context.Context, id league.League) (*Leagues, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.lolesports.com/api/v1/leagues?id=%d", id), nil)
	if err != nil {
		return nil, err
	}
	var res Leagues
	_, err = c.doJSON(ctx, req, &res)
	return &res, err
}

func (c Client) GetLeaguesBySlug(ctx context.Context, slug string) (*Leagues, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.lolesports.com/api/v1/leagues?slug=%s", slug), nil)
	if err != nil {
		return nil, err
	}
	var res Leagues
	_, err = c.doJSON(ctx, req, &res)
	return &res, err
}
