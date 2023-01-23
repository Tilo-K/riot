// Package region defines region constants.
package v5region

import "fmt"

// Region represents a Riot server region. Only constants defined in this
// package are valid inputs for the client.
type V5Region string

const (
	Americas = "AMERICAS"
	Asia = "ASIA"
	Europe = "EUROPE"
	Sea = "SEA"
)

// All returns all supported regions.
func All() []V5Region {
	return []V5Region{
		Americas,
		Asia,
		Europe,
		Sea,
	}
}

// Host returns the full hostname corresponding to the region. This function
// panics if an invalid region is used.
func (r V5Region) Host() string {
	switch r {
	case Americas:
		return "https://americas.api.riotgames.com"
	case Asia:
		return "https://asia.api.riotgames.com"
	case Europe:
		return "https://europe.api.riotgames.com"
	case Sea:
		return "https://sea.api.riotgames.com"
	default:
		panic(fmt.Sprintf("region %s does not have a configured host", r))
	}
}
