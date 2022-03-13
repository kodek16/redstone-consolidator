package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PriceFeedList: []PriceFeed{
      PriceFeed {
        Name: "kodek-test",
        BirthTime: "2022-03-13T00:00:00Z",
        Interval: "10m",
        Providers: []string{ "0xabc" },
        Assets: []string{ "USDUSD" },
      },
    },
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in priceFeed
	priceFeedIndexMap := make(map[string]struct{})

	for _, elem := range gs.PriceFeedList {
		index := string(PriceFeedKey(elem.Name))
		if _, ok := priceFeedIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for priceFeed")
		}
		priceFeedIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
