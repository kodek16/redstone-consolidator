package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PriceFeedKeyPrefix is the prefix to retrieve all PriceFeed
	PriceFeedKeyPrefix = "PriceFeed/value/"
)

// PriceFeedKey returns the store key to retrieve a PriceFeed from the index fields
func PriceFeedKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
