package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator/types"
)

// SetPriceFeed set a specific priceFeed in the store from its index
func (k Keeper) SetPriceFeed(ctx sdk.Context, priceFeed types.PriceFeed) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceFeedKeyPrefix))
	b := k.cdc.MustMarshal(&priceFeed)
	store.Set(types.PriceFeedKey(
		priceFeed.Name,
	), b)
}

// GetPriceFeed returns a priceFeed from its index
func (k Keeper) GetPriceFeed(
	ctx sdk.Context,
	name string,

) (val types.PriceFeed, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceFeedKeyPrefix))

	b := store.Get(types.PriceFeedKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePriceFeed removes a priceFeed from the store
func (k Keeper) RemovePriceFeed(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceFeedKeyPrefix))
	store.Delete(types.PriceFeedKey(
		name,
	))
}

// GetAllPriceFeed returns all priceFeed
func (k Keeper) GetAllPriceFeed(ctx sdk.Context) (list []types.PriceFeed) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceFeedKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PriceFeed
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
