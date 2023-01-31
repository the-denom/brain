package keeper

import (
	"github.com/cdbo/brain/x/membership/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetMember(ctx sdk.Context, address sdk.AccAddress) (member types.Member, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberKey))
	key := types.MakeMemberAddressKey(address)

	k.Logger(ctx).Debug("GetMember", "address", address)

	var b []byte
	if b = store.Get(key); b == nil {
		return member, false
	}

	k.cdc.Unmarshal(b, &member)
	return member, true
}

func (k Keeper) SetMember(ctx sdk.Context, address sdk.AccAddress, newMember types.Member) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberKey))
	memberCount := k.GetMemberCount(ctx)
	key := types.MakeMemberAddressKey(address)

	// Marshal and Set
	memberData := k.cdc.MustMarshal(&newMember)
	store.Set(key, memberData)

	// Bump member count
	k.SetMemberCount(ctx, memberCount+1)
}

func (k Keeper) IsMember(ctx sdk.Context, address sdk.AccAddress) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberKey))
	key := types.MakeMemberAddressKey(address)
	return store.Has(key)
}

func (k Keeper) GetMemberCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberCountKey))
	byteKey := []byte(types.MemberKey)
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
