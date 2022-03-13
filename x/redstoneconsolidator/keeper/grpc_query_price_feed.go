package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PriceFeedAll(c context.Context, req *types.QueryAllPriceFeedRequest) (*types.QueryAllPriceFeedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var priceFeeds []types.PriceFeed
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	priceFeedStore := prefix.NewStore(store, types.KeyPrefix(types.PriceFeedKeyPrefix))

	pageRes, err := query.Paginate(priceFeedStore, req.Pagination, func(key []byte, value []byte) error {
		var priceFeed types.PriceFeed
		if err := k.cdc.Unmarshal(value, &priceFeed); err != nil {
			return err
		}

		priceFeeds = append(priceFeeds, priceFeed)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPriceFeedResponse{PriceFeed: priceFeeds, Pagination: pageRes}, nil
}

func (k Keeper) PriceFeed(c context.Context, req *types.QueryGetPriceFeedRequest) (*types.QueryGetPriceFeedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPriceFeed(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetPriceFeedResponse{PriceFeed: val}, nil
}
