package keeper

import (
	"github.com/cdbo/brain/x/membership/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetMember(ctx sdk.Context, address sdk.AccAddress, newMember types.Member) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberKey))
	// TODO: remove - checkpoint for demo
	_ = store
	return 0
}

func (k Keeper) GetMemberCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberCountKey))
	byteKey := []byte(types.MemberCountKey)
	bz := store.Get(byteKey)

	// Nil result means zero members
	if bz == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetMemberCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberCountKey))
	bz := sdk.Uint64ToBigEndian(count)
	store.Set([]byte(types.MemberCountKey), bz)
}
